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
	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"
	"os"
)

// docCmd represents the doc command
func NewDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "doc",
		Short:        "generate doc",
		Long:         `generate doc`,
		RunE:         buildDoc,
		SilenceUsage: true,
	}
	return cmd
}

func buildDoc(cmd *cobra.Command, args []string) error {
	r := routes()
	doc := docgen.JSONRoutesDoc(r.(chi.Router))
	file, err := os.OpenFile("api_doc.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(doc); err != nil {
		return err
	}
	md := docgen.MarkdownRoutesDoc(r.(chi.Router), docgen.MarkdownOpts{
		ProjectPath: "github.com/go-chi/chi",
		Intro:       "Welcome to the stone api docs.",
	})
	mdFile, err := os.OpenFile("api_doc.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		return err
	}
	defer mdFile.Close()
	_, err = mdFile.WriteString(md)
	return err
}
