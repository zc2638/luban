/**
 * Created by zc on 2020/7/26.
 */
package resource

import (
	"github.com/go-chi/chi"
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/pkg/errs"
	"luban/service"
	"net/http"
)

func VersionList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Resource().VersionList(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		result := make([]response.VersionResultItem, 0, len(list))
		for _, v := range list {
			result = append(result, response.VersionResultItem{
				VersionID:  v.VersionID,
				ResourceID: v.ResourceID,
				Version:    v.Version,
				Format:     v.Format,
				Desc:       v.Desc,
				Content:    v.Content,
				Timestamp: response.Timestamp{
					CreatedTS: v.CreatedAt.Unix(),
					UpdatedTS: v.UpdatedAt.Unix(),
				},
			})
		}
		ctr.OK(w, result)
	}
}

func VersionInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		version := chi.URLParam(r, "version")
		info, err := service.New().Resource().VersionFind(r.Context(), version)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, response.VersionResultItem{
			VersionID:  info.VersionID,
			ResourceID: info.ResourceID,
			Version:    info.Version,
			Format:     info.Format,
			Desc:       info.Desc,
			Content:    info.Content,
			Timestamp: response.Timestamp{
				CreatedTS: info.CreatedAt.Unix(),
				UpdatedTS: info.UpdatedAt.Unix(),
			},
		})
	}
}

func VersionCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.ResourceVersionParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Version) {
			ctr.BadRequest(w, errs.ErrInvalidResourceVersion.With(compile.NameError))
			return
		}
		version := &data.Version{
			Version: params.Version,
			Desc:    params.Desc,
		}
		if err := service.New().Resource().VersionCreate(r.Context(), version); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func VersionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		version := chi.URLParam(r, "version")
		if err := service.New().Resource().VersionDelete(r.Context(), version); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
