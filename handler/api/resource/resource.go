/**
 * Created by zc on 2020/7/19.
 */
package resource

import (
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/pkg/errs"
	"luban/service"
	"net/http"
)

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Resource().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		result := make([]response.ResourceResultItem, 0, len(list))
		for _, v := range list {
			result = append(result, response.ResourceResultItem{
				ResourceID: v.ResourceID,
				SpaceID:    v.SpaceID,
				Name:       v.Name,
				Desc:       v.Desc,
				Format:     v.Format,
				Content:    v.Content,
				Label:      v.Label,
				Timestamp: response.Timestamp{
					CreatedTS: v.CreatedAt.Unix(),
					UpdatedTS: v.UpdatedAt.Unix(),
				},
			})
		}
		ctr.OK(w, result)
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := service.New().Resource().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, response.ResourceResultItem{
			ResourceID: info.ResourceID,
			SpaceID:    info.SpaceID,
			Name:       info.Name,
			Desc:       info.Desc,
			Format:     info.Format,
			Content:    info.Content,
			Label:      info.Label,
			Timestamp: response.Timestamp{
				CreatedTS: info.CreatedAt.Unix(),
				UpdatedTS: info.UpdatedAt.Unix(),
			},
		})
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.ResourceParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidResource.With(compile.NameError))
			return
		}
		resource := &data.Resource{
			Name:    params.Name,
			Desc:    params.Desc,
			Format:  params.Format,
			Content: params.Content,
			Label:   params.Label,
		}
		if err := service.New().Resource().Create(r.Context(), resource); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.ResourceParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidResource.With(compile.NameError))
			return
		}
		resource := &data.Resource{
			Name:    params.Name,
			Desc:    params.Desc,
			Format:  params.Format,
			Content: params.Content,
			Label:   params.Label,
		}
		if err := service.New().Resource().Update(r.Context(), resource); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Resource().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
