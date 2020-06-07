/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"stone/pkg/ctr"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer, middleware.Logger, cors.AllowAll().Handler)
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctr.Str(w, "Hello World!")
	})
	return mux
}