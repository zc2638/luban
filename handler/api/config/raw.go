/**
 * Created by zc on 2020/7/19.
 */
package config

import (
	"github.com/go-chi/chi"
	"luban/pkg/ctr"
	"luban/service"
	"net/http"
)

func Raw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		space := chi.URLParam(r, "space")
		config := chi.URLParam(r, "config")
		raw, err := service.New().Config().Raw(r.Context(), username, space, config)
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
		config := chi.URLParam(r, "config")
		version := chi.URLParam(r, "version")
		raw, err := service.New().Config().VersionRaw(r.Context(), username, space, config, version)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Bytes(w, raw)
	}
}
