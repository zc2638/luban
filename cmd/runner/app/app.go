/**
 * Created by zc on 2020/8/2.
 */
package app

import (
	"context"
	"docker.io/go-docker"
	"fmt"
	dockerengine "github.com/drone/drone-runtime/engine/docker"
	"github.com/drone/drone-runtime/runtime"
	"github.com/drone/drone-yaml/yaml"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"luban/cmd/internal/env"
	"luban/pkg/api/data"
	"luban/pkg/drone"
	"luban/pkg/errs"
	"luban/service/client"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "luban",
		Short: "luban runner",
		Long:  `Luban Runner.`,
		Run:   run,
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", env.Config(), "config file (default is $HOME/config.yaml)")
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	//config, err := global.ParseConfig(cfgFile)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	ctx, cancel := context.WithCancel(context.Background())
	if err := listen(ctx); err != nil {
		fmt.Println("program terminated")
		return
	}
	exit(ctx, cancel)
}

func exit(ctx context.Context, cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)
	select {
	case <-ctx.Done():
	case <-c:
		println("interrupt received, terminating process")
		cancel()
	}
}

func listen(ctx context.Context) error {
	var group errgroup.Group
	// TODO add num
	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				exec()
			}
		}
	})
	return group.Wait()
}

func exec() {
	fmt.Println("runner: polling queue start")
	// TODO 查询队列
	cli := client.New("http://localhost:8080")
	task, err := cli.Request(context.Background())
	if err != nil {
		fmt.Println(errs.Error("runner: cannot get queue item").With(err).Error())
		return
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	// TODO 异步监听取消信号
	// TODO 执行
	if err := start(task); err != nil {
		fmt.Println(err)
	}
}

func start(task *data.Task) error {
	// TODO 获取触发器流程配置
	// TODO 获取配置内容
	// TODO 转换到engine可执行结构体
	// TODO 自动构建容器文件
	// TODO 构建hook处理运行，实时返回请求server
	// TODO 构建runner并执行

	var configData = []byte(`a: 1`)
	var pipeline = &yaml.Pipeline{
		Kind: "pipeline",
		Type: "docker",
		Name: "test",
		Steps: []*yaml.Container{
			{
				Name:  "test",
				Image: "alpine",
				Commands: []string{
					`echo "Hello World!"`,
					`cat /work/.luban-config`,
				},
			},
		},
		Clone: yaml.Clone{Disable: true},
	}
	spec, err := drone.Compile(pipeline)
	if err != nil {
		return err
	}
	drone.Preset(spec, configData)
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	eng := dockerengine.New(dockerClient)
	hooks := &runtime.Hook{
		BeforeEach: func(s *runtime.State) error {
			return nil
		},
		AfterEach: func(s *runtime.State) error {
			return nil
		},
		GotLine: func(s *runtime.State, line *runtime.Line) error {
			fmt.Println(line.Number, line.Message)
			return nil
		},
	}
	runner := runtime.New(
		runtime.WithEngine(eng),
		runtime.WithConfig(spec),
		runtime.WithHooks(hooks),
	)

	timeout, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	fmt.Println("runner: start execution")
	err = runner.Run(timeout)
	if err != nil && err != runtime.ErrInterrupt {
		fmt.Println("runner: execution failed", err)
		return err
	}
	fmt.Println("runner: execution complete")
	return nil
}