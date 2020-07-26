/**
 * Created by zc on 2020/7/19.
 */
package config

import (
	"luban/pkg/api"
	"luban/pkg/api/store"
	"luban/pkg/ctr"
	"luban/service"
	"net/http"
)

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Config().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := service.New().Config().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, info)
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.ConfigParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		config := store.Config{
			Name:    params.Name,
			Desc:    params.Desc,
			Format:  params.Format,
			Content: params.Content,
		}
		if err := service.New().Config().Create(r.Context(), &config); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.ConfigParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		config := store.Config{
			Name:    params.Name,
			Desc:    params.Desc,
			Format:  params.Format,
			Content: params.Content,
		}
		if err := service.New().Config().Update(r.Context(), &config); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Config().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
