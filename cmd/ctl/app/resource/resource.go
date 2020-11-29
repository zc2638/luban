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
package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"luban/handler/api/resource"
	"luban/pkg/api/response"
	"luban/pkg/api/stdout"
	"luban/pkg/printer"
	"net/http"
)

var kind string

// NewCmd represents the resource command
// lubanctl resource list -k pipeline
// lubanctl resource [list/get/delete/version]
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource",
		Short: "resource operation",
		Long:  `resource operation`,
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:          "get",
			Short:        "resource operation",
			Long:         `resource operation`,
			RunE:         get,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "delete",
			Short:        "resource operation",
			Long:         `resource operation`,
			RunE:         del,
			SilenceUsage: true,
		},
		&cobra.Command{
			Use:          "version",
			Short:        "resource version operation",
			Long:         `resource version operation`,
			RunE:         version,
			SilenceUsage: true,
		},
	)
	listCmd := &cobra.Command{
		Use:          "list",
		Short:        "resource operation",
		Long:         `resource operation`,
		RunE:         list,
		SilenceUsage: true,
	}
	listCmd.Flags().StringVarP(&kind, "kind", "k", "", "resource kind: file|pipeline")
	cmd.AddCommand(listCmd)
	return cmd
}

var client = resty.New().SetHostURL("http://localhost:2639/v1")

func list(cmd *cobra.Command, args []string) error {
	resp, err := client.R().
		SetQueryParam("kind", kind).Get(resource.Path)
	if err != nil {
		return err
	}
	if resp.RawResponse.StatusCode != http.StatusOK {
		return errors.New(resp.String())
	}
	var list []stdout.ResourceItem
	if err := json.Unmarshal(resp.Body(), &list); err != nil {
		return err
	}
	format := "%s\t%s\t%s\t%s\n"
	w := printer.New()
	fmt.Fprintf(w, format, "NAME", "KIND", "FORMAT", "DESC")
	for _, v := range list {
		fmt.Fprintf(w, format, v.Name, v.Kind, v.Format, v.Desc)
	}
	return w.Flush()
}

func get(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resource name not found")
	}
	resp, err := client.R().
		SetPathParams(map[string]string{
			"resource": args[0],
		}).Get(resource.PathName)
	if err != nil {
		return err
	}
	if resp.RawResponse.StatusCode != http.StatusOK {
		return errors.New(resp.String())
	}
	var resource stdout.ResourceItem
	if err := json.Unmarshal(resp.Body(), &resource); err != nil {
		return err
	}
	w := printer.New()
	fmt.Fprintln(w, resource.Data)
	return nil
}

func del(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resource name not found")
	}
	resp, err := client.R().
		SetPathParams(map[string]string{
			"resource": args[0],
		}).Delete(resource.PathName)
	if err != nil {
		return err
	}
	if resp.RawResponse.StatusCode != http.StatusOK {
		return errors.New(resp.String())
	}
	return nil
}

func version(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resource name not found")
	}
	resp, err := client.R().
		SetPathParams(map[string]string{
			"resource": args[0],
		}).Get(resource.PathVersion)
	if err != nil {
		return err
	}
	if resp.RawResponse.StatusCode != http.StatusOK {
		return errors.New(string(resp.Body()))
	}
	var list []response.VersionResultItem
	if err := json.Unmarshal(resp.Body(), &list); err != nil {
		return errors.New(resp.String())
	}
	format := "%s\t%s\t%s\t%s\n"
	w := printer.New()
	fmt.Fprintf(w, format, "VERSION", "KIND", "FORMAT", "DESC")
	for _, v := range list {
		fmt.Fprintf(w, format, v.Version, v.Kind, v.Format, v.Desc)
	}
	return w.Flush()
}
