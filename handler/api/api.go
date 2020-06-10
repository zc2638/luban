/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"stone/handler/api/auth"
	"stone/pkg/ctr"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer, middleware.Logger, cors.AllowAll().Handler)
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctr.Str(w, "Hello Stone!")
	})
	mux.Route("/auth", authRoute)
	return mux
}

func authRoute(mux chi.Router) {
	mux.Post("/register", auth.Register())
	mux.Post("/login", auth.Login())
}
