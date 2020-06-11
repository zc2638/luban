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
	"stone/handler/api/space"
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

// v1 handle v1 version routing
func v1(r chi.Router) {
	r.Route("/auth", authRoute)
	r.Group(func(r chi.Router) {
		r.Use(JwtAuth)
		r.Route("/user", userRoute)
		r.Route("/space", spaceRoute)
	})
}

// authRoute handle auth routing related
func authRoute(r chi.Router) {
	r.Post("/register", auth.Register())
	r.Post("/login", auth.Login())
}

// userRoute handle user routing related
func userRoute(r chi.Router) {
	r.Get("/", auth.Info())
}

// spaceRoute handle space routing related
func spaceRoute(r chi.Router) {
	r.Get("/", space.List())
	r.Post("/", space.Create())
	r.Get("/{id}", space.Find())
	r.Put("/{id}", space.Update())
	r.Delete("/{id}", space.Delete())
}
