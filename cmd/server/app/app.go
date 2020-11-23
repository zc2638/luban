/**
 * Created by zc on 2020/8/2.
 */
package app

import (
	"fmt"
	"github.com/pkgms/go/server"
	"github.com/spf13/cobra"
	"luban/cmd/internal/env"
	"luban/cmd/internal/route"
	"luban/global"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "luban",
		Short: "luban server",
		Long:  `Luban Server.`,
		RunE:  run,
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", env.Config(), "config file (default is $HOME/config.yaml)")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := global.ParseConfig(cfgFile)
	if err != nil {
		return err
	}
	if err := global.InitConfig(cfg); err != nil {
		return err
	}
	s := server.New(&cfg.Server)
	s.Handler = route.New()
	fmt.Println("Listen on", s.Addr)
	return s.Run()
}
