/**
 * Created by zc on 2020/6/11.
 */
package space

import (
	"github.com/pkgms/go/ctr"
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/errs"
	"luban/pkg/store"
	"luban/pkg/wrap"
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
		result := make([]response.SpaceResultItem, 0, len(list))
		for _, v := range list {
			result = append(result, response.SpaceResultItem{
				SpaceID: v.SpaceID,
				Name:    v.Name,
				Timestamp: response.Timestamp{
					CreatedTS: v.CreatedAt.Unix(),
					UpdatedTS: v.UpdatedAt.Unix(),
				},
			})
		}
		ctr.OK(w, result)
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.SpaceParams
		if err := wrap.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidSpace.With(compile.NameError))
			return
		}
		space := &store.Space{Name: params.Name}
		if err := service.New().Space().Create(r.Context(), space); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.SpaceParams
		if err := wrap.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidSpace.With(compile.NameError))
			return
		}
		space := &store.Space{Name: params.Name}
		if err := service.New().Space().Update(r.Context(), space); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Space().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
