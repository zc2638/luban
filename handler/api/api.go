/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"luban/handler/api/auth"
	"luban/handler/api/space"
	"net/http"
)

// New handle v1 version routing
func New() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer, middleware.Logger, cors.AllowAll().Handler)
	mux.Route("/auth", authRoute)
	mux.Group(func(r chi.Router) {
		r.Use(JwtAuth)
		r.Route("/user", userRoute)
		r.Route("/space", spaceRoute)
	})
	return mux
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
