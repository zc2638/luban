/**
 * Created by zc on 2020/8/10.
**/
package pipeline

import (
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"luban/pkg/api/request"
	"luban/pkg/store"
	"luban/pkg/wrapper"
	"luban/service"
	"net/http"
)

// Route handle pipeline routing related
func Route(doc *swagger.API) {
	const tag = "pipeline"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "流水线管理",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodGet, "/pipeline",
			endpoint.Handler(list()),
			endpoint.Summary("流水线列表"),
			endpoint.ResponseSuccess(endpoint.Schema([]store.Pipeline{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodGet, "/pipeline/{pipeline}",
			endpoint.Handler(info()),
			endpoint.Summary("流水线详情"),
			endpoint.Path("pipeline", "string", "流水线标识", true),
			endpoint.ResponseSuccess(endpoint.Schema(store.Pipeline{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPost, "/pipeline",
			endpoint.Handler(create()),
			endpoint.Summary("流水线创建"),
			endpoint.Body(request.PipelineParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPut, "/pipeline/{pipeline}",
			endpoint.Handler(update()),
			endpoint.Summary("流水线更新"),
			endpoint.Path("pipeline", "string", "流水线标识", true),
			endpoint.Body(request.PipelineParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodDelete, "/pipeline/{pipeline}",
			endpoint.Handler(del()),
			endpoint.Summary("流水线删除"),
			endpoint.Path("pipeline", "string", "流水线标识", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Pipeline().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := service.New().Pipeline().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, info)
	}
}

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.PipelineParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
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

func update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.PipelineParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
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

func del() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := service.New().Pipeline().Delete(r.Context()); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
