/**
 * Created by zc on 2020/7/19.
 */
package resource

import (
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/errs"
	"luban/pkg/store"
	"luban/pkg/wrapper"
	"luban/service"
	"net/http"
)

const (
	Path     = "/resource"
	PathName = "/resource/{resource}"
)

// Route handle resource routing related
func Route(doc *swagger.API) {
	const tag = "resource"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "资源管理",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodGet, Path,
			endpoint.Handler(list()),
			endpoint.Summary("资源列表"),
			endpoint.ResponseSuccess(endpoint.Schema([]response.ResourceResultItem{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodGet, PathName,
			endpoint.Handler(info()),
			endpoint.Summary("资源详情"),
			endpoint.Path("resource", "string", "资源名称", true),
			endpoint.ResponseSuccess(endpoint.Schema(response.ResourceResultItem{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPost, Path,
			endpoint.Handler(create()),
			endpoint.Summary("资源创建"),
			endpoint.Body(request.ResourceParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPut, PathName,
			endpoint.Handler(update()),
			endpoint.Summary("资源更新"),
			endpoint.Path("resource", "string", "资源名称", true),
			endpoint.Body(request.ResourceParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodDelete, PathName,
			endpoint.Handler(del()),
			endpoint.Summary("资源删除"),
			endpoint.Path("resource", "string", "资源名称", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kind := r.URL.Query().Get("kind")
		list, err := service.New().Resource().List(r.Context(), kind)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		result := make([]response.ResourceResultItem, 0, len(list))
		for _, v := range list {
			result = append(result, response.ResourceResultItem{
				Name:   v.Name,
				Desc:   v.Desc,
				Kind:   v.Kind,
				Format: v.Format,
				Label:  v.Label,
				Timestamp: response.Timestamp{
					CreatedTS: v.CreatedAt.Unix(),
					UpdatedTS: v.UpdatedAt.Unix(),
				},
			})
		}
		ctr.OK(w, result)
	}
}

func info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "resource")
		info, err := service.New().Resource().Find(r.Context(), name)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, response.ResourceResultItem{
			Name:   info.Name,
			Desc:   info.Desc,
			Kind:   info.Kind,
			Format: info.Format,
			Data:   info.Data,
			Label:  info.Label,
			Timestamp: response.Timestamp{
				CreatedTS: info.CreatedAt.Unix(),
				UpdatedTS: info.UpdatedAt.Unix(),
			},
		})
	}
}

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.ResourceParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidResource.With(compile.NameError))
			return
		}
		resource := &store.Resource{
			Name:   params.Name,
			Desc:   params.Desc,
			Format: params.Format,
			Data:   params.Content,
			Label:  params.Label,
		}
		if err := service.New().Resource().Create(r.Context(), resource); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.ResourceParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Name().MatchString(params.Name) {
			ctr.BadRequest(w, errs.ErrInvalidResource.With(compile.NameError))
			return
		}
		resource := &store.Resource{
			Name:   params.Name,
			Desc:   params.Desc,
			Format: params.Format,
			Data:   params.Content,
			Label:  params.Label,
		}
		if err := service.New().Resource().Update(r.Context(), resource); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func del() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "resource")
		if err := service.New().Resource().Delete(r.Context(), name); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
