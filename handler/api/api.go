/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"luban/handler/api/auth"
	"luban/handler/api/config"
	"luban/handler/api/space"
	"luban/pkg/ctr"
	"net/http"
)

// New handle v1 version routing
func New() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer, middleware.Logger, cors.AllowAll().Handler)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctr.Str(w, "Hello Luban!")
	}))
	mux.Route("/auth", authRoute)
	mux.Route("/raw/{username}/{space}/{config}", func(r chi.Router) {
		r.Get("/", config.Raw())
		r.Get("/{version}", config.VersionRaw())
	})
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
	r.Route("/{space}", func(cr chi.Router) {
		cr.Use(SpaceAuth)
		cr.Put("/", space.Update())
		cr.Delete("/", space.Delete())
		cr.Route("/config", configRoute)
	})
}

// configRoute handle config routing related
func configRoute(r chi.Router) {
	r.Get("/", config.List())
	r.Post("/", config.Create())
	r.Route("/{config}", func(cr chi.Router) {
		cr.Use(ConfigAuth)
		cr.Get("/", config.Info())
		cr.Put("/", config.Update())
		cr.Delete("/", config.Delete())
		cr.Route("/version", configVersionRoute)
	})
}

// configVersionRoute handle config version routing related
func configVersionRoute(r chi.Router) {
	r.Get("/", config.VersionList())
	r.Post("/", config.VersionCreate())
	r.Route("/{name}", func(cr chi.Router) {
		cr.Get("/", config.VersionFind())
		cr.Delete("/", config.VersionDelete())
	})
}
