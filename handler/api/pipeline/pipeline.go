/**
 * Created by zc on 2020/8/10.
**/
package pipeline

import (
	"github.com/pkgms/go/ctr"
	"luban/pkg/api/request"
	"luban/pkg/store"
	"luban/pkg/wrap"
	"luban/service"
	"net/http"
)

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Pipeline().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := service.New().Pipeline().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, info)
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.PipelineParams
		if err := wrap.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		pipeline := &store.Pipeline{
			ResourceID: params.ResourceID,
			Name:       params.Name,
			Spec:       params.Spec,
		}
		if err := service.New().Pipeline().Create(r.Context(), pipeline); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.PipelineParams
		if err := wrap.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		pipeline := &store.Pipeline{
			ResourceID: params.ResourceID,
			Name:       params.Name,
			Spec:       params.Spec,
		}
		if err := service.New().Pipeline().Update(r.Context(), pipeline); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Pipeline().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
