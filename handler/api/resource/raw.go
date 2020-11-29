/**
 * Created by zc on 2020/7/19.
 */
package resource

import (
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"luban/service"
	"net/http"
)

// RawRoute handle raw routing related
func RawRoute(doc *swagger.API) {
	const tag = "raw"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "webhook",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodGet, "/raw/{username}/{resource}",
			endpoint.Handler(raw()),
			endpoint.Summary("资源信息获取"),
			endpoint.Path("username", "string", "用户名", true),
			endpoint.Path("resource", "string", "资源名", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodGet, "/raw/{username}/{resource}/{version}",
			endpoint.Handler(versionRaw()),
			endpoint.Summary("资源版本信息获取"),
			endpoint.Path("username", "string", "用户名", true),
			endpoint.Path("resource", "string", "资源名", true),
			endpoint.Path("version", "string", "资源版本", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func raw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		resource := chi.URLParam(r, "resource")
		raw, err := service.New().Resource().Raw(r.Context(), username, resource)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Bytes(w, raw)
	}
}

func versionRaw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		resource := chi.URLParam(r, "resource")
		version := chi.URLParam(r, "version")
		raw, err := service.New().Resource().VersionRaw(r.Context(), username, resource, version)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Bytes(w, raw)
	}
}
