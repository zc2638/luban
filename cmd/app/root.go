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
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"stone/pkg/global"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stone",
		Short: "Stone service",
		Long:  `Stone service.`,
	}
	cmd.AddCommand(NewServerCmd(), NewMigrateCmd(), NewConfigCmd())
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", global.DefaultPath, "config file (default is $HOME/config.yaml)")
	return cmd
}

func ParseConfig() (*global.Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.yaml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	cfg := global.Environ()
	err := viper.Unmarshal(cfg)
	return cfg, err
}
