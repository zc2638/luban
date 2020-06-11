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
	mux.Route("/v1", v1)
	return mux
}

func v1(r chi.Router) {
	r.Route("/auth", authRoute)
	r.Group(func(r chi.Router) {
		r.Use(JwtAuth)
		r.Route("/user", userRoute)
	})
}

func authRoute(r chi.Router) {
	r.Post("/register", auth.Register())
	r.Post("/login", auth.Login())
}

func userRoute(r chi.Router) {
	r.Get("/", auth.Info())
}
