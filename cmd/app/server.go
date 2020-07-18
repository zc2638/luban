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
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"luban/global"
	"luban/handler/api"
	"luban/pkg/ctr"
	"luban/pkg/server"
	"net/http"
)

// serverCmd represents the server command
func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Short:        "Run server",
		Long:         `Run server`,
		RunE:         startServer,
		SilenceUsage: true,
	}
	return cmd
}

func startServer(cmd *cobra.Command, args []string) error {
	cfg, err := ParseConfig()
	if err != nil {
		return err
	}
	if err := global.InitConfig(cfg); err != nil {
		return err
	}
	s := server.New(&cfg.Server)
	s.Handler = routes()
	fmt.Println("Listen on", s.Addr)
	return s.Run()
}

func routes() http.Handler {
	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctr.Str(w, "Hello Luban!")
	})
	mux.Mount("/v1", api.New())
	return mux
}