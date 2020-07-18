/**
 * Created by zc on 2020/6/11.
 */
package space

import (
	"github.com/go-chi/chi"
	"luban/pkg/api"
	"luban/pkg/api/store"
	"luban/pkg/ctr"
	"luban/service"
	"net/http"
)

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Space().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		space, err := service.New().Space().Find(r.Context(), id)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, space)
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.SpaceParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		space := store.Space{
			Title:       params.Title,
			Description: params.Description,
		}
		if err := service.New().Space().Create(r.Context(), &space); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.SpaceParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		id := chi.URLParam(r, "id")
		space := store.Space{
			SID:         id,
			Title:       params.Title,
			Description: params.Description,
		}
		if err := service.New().Space().Update(r.Context(), &space); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := service.New().Space().Delete(r.Context(), id); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
