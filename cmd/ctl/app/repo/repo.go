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
package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/drone/drone/core"
	"github.com/spf13/cobra"
	"github.com/zc2638/drone-control/client"
	"github.com/zc2638/drone-control/handler/api"
	"io/ioutil"
	"luban/pkg/printer"
	"luban/pkg/util"
	"strconv"
	"strings"
)

var file string

// NewCmd represents the pipeline command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "repo operation",
		Long:  `repo operation`,
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:          "list",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         list,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "get",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         get,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "delete",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         del,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "run",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         run,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "record",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         record,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "stage",
			Short:        "repo operation",
			Long:         `repo operation`,
			RunE:         stage,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "logs",
			Short:        "task operation",
			Long:         `task operation`,
			RunE:         logs,
			SilenceUsage: true,
		},
	)
	applyCmd := &cobra.Command{
		Use:          "apply",
		Short:        "repo operation",
		Long:         `repo operation`,
		RunE:         apply,
		SilenceUsage: true,
	}
	applyCmd.Flags().StringVarP(&file, "file", "f", "", "pipeline spec file path")
	cmd.AddCommand(applyCmd)
	return cmd
}

const host = "http://localhost:2639/api"

func list(cmd *cobra.Command, args []string) error {
	c := client.New(client.Config{Host: host})
	list, err := c.Repo("default", "").List()
	if err != nil {
		return err
	}
	t := printer.NewTab("NAME", "NAMESPACE", "CREATED")
	for _, v := range list {
		t.Add(v.Name, v.Namespace, util.SecTimestampToDateTime(v.Created))
	}
	t.Print()
	return nil
}

func get(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("repo name not found")
	}
	c := client.New(client.Config{Host: host})
	data, err := c.Repo("default", args[0]).Info()
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("repo spec not found")
	}
	fmt.Println(data.Data)
	return nil
}

func apply(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("repo name not found")
	}
	file = strings.TrimSpace(file)
	if file == "" {
		return errors.New("file path is empty")
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	c := client.New(client.Config{Host: host})
	return c.Repo("default", "").Apply(&api.RepoParams{
		Namespace: "default",
		Name:      args[0],
		Data:      string(b),
	})
}

func del(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("repo name is not found")
	}
	c := client.New(client.Config{Host: host})
	return c.Repo("default", args[0]).Delete()
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("repo name not found")
	}
	c := client.New(client.Config{Host: host})
	return c.Build("default", args[0]).Run()
}

func record(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("repo name is not found")
	}
	c := client.New(client.Config{Host: host})
	list, err := c.Build("default", args[0]).List()
	if err != nil {
		return err
	}
	t := printer.NewTab("NUMBER", "STATUS", "STARTTED", "FINISHED")
	for _, v := range list {
		t.Add(
			strconv.FormatInt(v.Number, 10),
			v.Status,
			util.SecTimestampToDateTime(v.Started),
			util.SecTimestampToDateTime(v.Finished),
		)
	}
	t.Print()
	return nil
}

func stage(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("params is not enough, [repo] [build_number]")
	}
	number, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("build number parse failed")
	}
	c := client.New(client.Config{Host: host})
	build, err := c.Build("default", args[0]).Info(number)
	if err != nil {
		return err
	}
	t := printer.NewTab()
	t.Add("Build Number:", strconv.FormatInt(build.Number, 10))
	t.Add("Build Status:", build.Status)
	t.Add("Build Created:", util.SecTimestampToDateTime(build.Created))
	t.Add("Build Started:", util.SecTimestampToDateTime(build.Started))
	t.Add("Build Finished:", util.SecTimestampToDateTime(build.Finished))
	t.Print()
	fmt.Println("========================================")
	for _, stage := range build.Stages {
		t.Add("Stage Number:", strconv.Itoa(stage.Number))
		t.Add("Stage Name:", stage.Name)
		t.Add("Stage Status:", stage.Status)
		t.Add("Stage Type:", stage.Type)
		t.Add("Stage RunBy:", stage.Machine)
		t.Add("Stage Created:", util.SecTimestampToDateTime(stage.Created))
		t.Add("Stage Started:", util.SecTimestampToDateTime(stage.Started))
		t.Add("Stage Stopped:", util.SecTimestampToDateTime(stage.Stopped))
		t.Print()
		fmt.Println()
	}
	return nil
}

func logs(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return errors.New("params is not enough")
	}
	build, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("build number parse failed")
	}
	stage, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("stage number parse failed")
	}
	step := -1
	if len(args) > 3 {
		step, err = strconv.Atoi(args[3])
		if err != nil {
			return errors.New("step number parse failed")
		}
	}

	c := client.New(client.Config{Host: host})
	b, err := c.Build("default", args[0]).Log(build, stage, step)
	if err != nil {
		return err
	}
	var data [][]core.Line
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	for k, v := range data {
		fmt.Printf("========== Step %d ==========\n", k+1)
		for _, line := range v {
			fmt.Printf(line.Message)
		}
		fmt.Println()
	}
	return nil
}
