/**
 * Created by zc on 2020/8/2.
 */
package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkgms/go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	control "github.com/zc2638/drone-control/global"
	"github.com/zc2638/drone-control/handler"
	"github.com/zc2638/drone-control/store"
	"luban/cmd/internal/env"
	"luban/cmd/internal/route"
	"luban/global"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "luban",
		Short:        "luban server",
		Long:         `Luban Server.`,
		RunE:         run,
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", env.Config(), "config file (default is $HOME/config.yaml)")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := global.ParseConfig(cfgFile)
	if err != nil {
		return err
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	if err := global.InitConfig(cfg); err != nil {
		return err
	}

	//go runControl(cfg.Control)
	s := server.New(&cfg.Server)
	s.Handler = route.New()
	return s.Run()
}

func runControl(cfg control.Config) error {
	if err := control.InitCfg(&cfg); err != nil {
		return err
	}
	if ok := control.GormDB().Migrator().HasTable(&store.ReposData{}); !ok {
		if err := control.GormDB().Migrator().CreateTable(&store.ReposData{}); err != nil {
			return err
		}
	}
	r := chi.NewRouter()
	r.Mount("/rpc/v2", handler.RPC())
	r.Mount("/api", handler.API())
	r.Mount("/static", handler.Static())
	s := server.New(&cfg.Server)
	s.Handler = r
	fmt.Println("Luban Control Listen on", s.Addr)
	return s.Run()
}
