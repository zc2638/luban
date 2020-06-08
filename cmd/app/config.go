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
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"stone/pkg/global"
	"stone/pkg/logger"
)

// configCmd represents the config command
func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "config operation",
		Long:  `config operation.`,
	}
	generateCmd := &cobra.Command{
		Use:          "new",
		Short:        "config generate operation",
		Long:         `config generate operation.`,
		RunE:         configureNew,
		SilenceUsage: true,
	}
	cmd.AddCommand(generateCmd)
	return cmd
}

func configureNew(cmd *cobra.Command, args []string) error {
	fmt.Println("Start to create config file: ", cfgFile)
	_, err := os.Stat(cfgFile)
	if err == nil {
		return errors.New(cfgFile + " is exist")
	}
	if os.IsExist(err) {
		return err
	}
	cfg := global.Environ()
	var buffer bytes.Buffer
	if err := yaml.NewEncoder(&buffer).Encode(cfg); err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(&buffer); err != nil {
		return err
	}
	fmt.Println("Loading default configuration succeeded")
	if err := viper.WriteConfigAs(cfgFile); err != nil {
		return err
	}
	fmt.Println("Completed")
	return nil
}
