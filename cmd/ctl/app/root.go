/*
Copyright © 2020 zc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"luban/cmd/ctl/app/repo"
	"luban/cmd/ctl/app/resource"
	"luban/cmd/internal/env"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	viper.AutomaticEnv()
	cmd := &cobra.Command{
		Use:   "lubanctl",
		Short: "luban service",
		Long:  `Luban Service.`,
	}
	cmd.AddCommand(
		NewMigrateCmd(),
		NewConfigCmd(),
		NewDocCmd(),
		NewUserCmd(),

		resource.NewCmd(),
		repo.NewCmd(),
	)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", env.Config(), "config file (default is $HOME/config.yaml)")
	return cmd
}
