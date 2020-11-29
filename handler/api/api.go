/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/drone-control/handler"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"luban/handler/api/auth"
	"luban/handler/api/pipeline"
	"luban/handler/api/resource"
	"luban/handler/api/space"
	"net/http"
)

var doc = swag.New(
	swag.Title("Luban API Doc"),
	swag.BasePath("/v1"),
)

func init() {
	doc.AddEndpoint(endpoint.New(
		http.MethodGet, "/",
		endpoint.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctr.Str(w, "Hello Luban!")
		})),
		endpoint.ResponseSuccess(),
	))
	doc.AddEndpointFunc(
		auth.Route,
		space.Route,
		resource.Route,
		resource.VersionRoute,
		resource.RawRoute,
		pipeline.Route,
	)
}

// New handle v1 version routing
func New() http.Handler {
	mux := chi.NewMux()
	mux.Use(
		middleware.Recoverer,
		middleware.Logger,
		cors.AllowAll().Handler,
		JwtAuth,
	)
	mux.Mount("/control", handler.API())
	for path, endpoints := range doc.Paths {
		mux.Handle(path, endpoints)
	}
	mux.Get("/swagger", doc.Handler(false))
	return mux
}
