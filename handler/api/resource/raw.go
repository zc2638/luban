/**
 * Created by zc on 2020/7/19.
 */
package resource

import (
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"luban/service"
	"net/http"
)

func Raw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		space := chi.URLParam(r, "space")
		resource := chi.URLParam(r, "resource")
		raw, err := service.New().Resource().Raw(r.Context(), username, space, resource)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Bytes(w, raw)
	}
}

func VersionRaw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		space := chi.URLParam(r, "space")
		resource := chi.URLParam(r, "resource")
		version := chi.URLParam(r, "version")
		raw, err := service.New().Resource().VersionRaw(r.Context(), username, space, resource, version)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Bytes(w, raw)
	}
}
