/**
 * Created by zc on 2020/6/11.
 */
package space

import (
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

// Route handle space routing related
func Route(doc *swagger.API) {
	const tag = "space"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "空间管理",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodGet, "/space",
			endpoint.Handler(list()),
			endpoint.Summary("空间列表"),
			endpoint.ResponseSuccess(endpoint.Schema([]response.SpaceResultItem{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPost, "/space",
			endpoint.Handler(create()),
			endpoint.Summary("空间创建"),
			endpoint.Body(request.SpaceParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPut, "/space/{space}",
			endpoint.Handler(update()),
			endpoint.Summary("空间更新"),
			endpoint.Path("space", "string", "空间标识", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodDelete, "/space/{space}",
			endpoint.Handler(del()),
			endpoint.Summary("空间删除"),
			endpoint.Path("space", "string", "空间标识", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func list() http.HandlerFunc {
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

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.SpaceParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
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

func update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.SpaceParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
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

func del() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Space().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
