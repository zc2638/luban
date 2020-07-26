/**
 * Created by zc on 2020/7/26.
 */
package config

import (
	"github.com/go-chi/chi"
	"luban/pkg/api"
	"luban/pkg/ctr"
	"luban/service"
	"net/http"
)

func VersionList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Config().VersionList(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func VersionFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		info, err := service.New().Config().VersionFind(r.Context(), name)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, string(info))
	}
}

func VersionCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.ConfigVersionParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if err := service.New().Config().VersionCreate(r.Context(), params.Version); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func VersionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if err := service.New().Config().VersionDelete(r.Context(), name); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}