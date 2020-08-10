/**
 * Created by zc on 2020/6/6.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"luban/handler/api/auth"
	"luban/handler/api/pipeline"
	"luban/handler/api/resource"
	"luban/handler/api/space"
	"luban/handler/api/task"
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
	mux.Route("/raw/{username}/{space}/{resource}", func(r chi.Router) {
		r.Get("/", resource.Raw())
		r.Get("/{version}", resource.VersionRaw())
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
		cr.Route("/resource", resourceRoute)
	})
}

// resourceRoute handle resource routing related
func resourceRoute(r chi.Router) {
	r.Get("/", resource.List())
	r.Post("/", resource.Create())
	r.Route("/{resource}", func(cr chi.Router) {
		cr.Use(ResourceAuth)
		cr.Get("/", resource.Info())
		cr.Put("/", resource.Update())
		cr.Delete("/", resource.Delete())
		cr.Route("/version", resourceVersionRoute)
	})
}

// resourceVersionRoute handle resource version routing related
func resourceVersionRoute(r chi.Router) {
	r.Get("/", resource.VersionList())
	r.Post("/", resource.VersionCreate())
	r.Route("/{version}", func(cr chi.Router) {
		cr.Get("/", resource.VersionInfo())
		cr.Delete("/", resource.VersionDelete())
	})
}

// pipelineRoute handle pipeline routing related
func pipelineRoute(r chi.Router) {
	r.Get("/", pipeline.List())
	r.Post("/", pipeline.Create())
	r.Route("/{pipeline}", func(cr chi.Router) {
		cr.Use(PipelineAuth)
		cr.Get("/", pipeline.Info())
		cr.Put("/", pipeline.Update())
		cr.Delete("/", pipeline.Delete())
		cr.Route("/task", taskRoute)
	})
}

// taskRoute handle task routing related
func taskRoute(r chi.Router) {
	r.Get("/", task.List())
	r.Post("/", task.Create())
	r.Route("/{task}", func(cr chi.Router) {
		cr.Use(TaskAuth)
		cr.Get("/", task.Info())
		cr.Route("/step", taskStepRoute)
	})
}

// taskStepRoute handle task step routing related
func taskStepRoute(r chi.Router) {
	r.Get("/", task.StepList())
	r.Put("/{step}", task.StepUpdate())
}
