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
	"github.com/spf13/cobra"
	"luban/global"
	"luban/global/database/migration"
)

// migrateCmd represents the migrate command
func NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "migrate",
		Short:        "database migration",
		Long:         `database migration.`,
		RunE:         migrate,
		SilenceUsage: true,
	}
	return cmd
}

func migrate(cmd *cobra.Command, args []string) error {
	fmt.Println("migration starting...")
	cfg, err := global.ParseConfig(cfgFile)
	if err != nil {
		return err
	}
	if err := migration.InitTable(&cfg.Database); err != nil {
		return err
	}
	fmt.Println("migration table successful")
	return nil
}
