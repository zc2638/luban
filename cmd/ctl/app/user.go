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
	"context"
	"github.com/spf13/cobra"
	"github.com/zc2638/gotool/utilx"
	"luban/pkg/compile"
	"luban/pkg/database/data"
	"luban/pkg/errs"
	"luban/service"
)

// NewUserCmd represents the user command
func NewUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user operation.",
		Long:  `user operation.`,
	}
	addCmd := &cobra.Command{
		Use:          "add",
		Short:        "add user operation",
		Long:         `add user operation.The first command argument is username and the second is password.`,
		RunE:         userAdd,
		SilenceUsage: true,
	}
	pwdResetCmd := &cobra.Command{
		Use:          "pwd-reset",
		Short:        "pwd reset operation",
		Long:         `pwd reset operation.Reset the password to the current user.`,
		RunE:         pwdReset,
		SilenceUsage: true,
	}
	cmd.AddCommand(addCmd)
	cmd.AddCommand(pwdResetCmd)
	return cmd
}

func userAdd(cmd *cobra.Command, args []string) error {
	argsLen := len(args)
	if argsLen < 2 {
		return errs.New("need username and password!")
	}
	if argsLen > 2 {
		return errs.New("unknown command argument: " + args[2])
	}
	username := args[0]
	if !compile.Username().MatchString(username) {
		return compile.UsernameError
	}
	password := args[1]
	return service.New().User().Create(context.Background(), &data.User{
		Username: username,
		Pwd:      password,
		Salt:     utilx.RandomStr(6),
	})
}

func pwdReset(cmd *cobra.Command, args []string) error {
	argsLen := len(args)
	if argsLen < 2 {
		return errs.New("need username and password!")
	}
	if argsLen > 2 {
		return errs.New("unknown command argument: " + args[2])
	}
	username := args[0]
	password := args[1]
	return service.New().User().PwdReset(context.Background(), username, password)
}
