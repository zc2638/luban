/**
 * Created by zc on 2020/7/26.
 */
package resource

import (
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"luban/pkg/api/response"
	"luban/service"
	"net/http"
)

const (
	PathVersion     = "/resource/{resource}/version"
	PathVersionName = "/resource/{resource}/version/{version}"
)

// VersionRoute handle resource version routing related
func VersionRoute(doc *swagger.API) {
	const tag = "resourceVersion"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "资源版本管理",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodGet, PathVersion,
			endpoint.Handler(versionList()),
			endpoint.Summary("资源版本列表"),
			endpoint.Path("resource", "resource", "资源名称", true),
			endpoint.ResponseSuccess(endpoint.Schema([]response.VersionResultItem{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodGet, PathVersionName,
			endpoint.Handler(versionInfo()),
			endpoint.Summary("资源版本详情"),
			endpoint.Path("resource", "resource", "资源名称", true),
			endpoint.Path("version", "string", "资源版本", true),
			endpoint.ResponseSuccess(endpoint.Schema(response.VersionResultItem{})),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodDelete, PathVersionName,
			endpoint.Handler(versionDelete()),
			endpoint.Summary("资源版本删除"),
			endpoint.Path("resource", "resource", "资源名称", true),
			endpoint.Path("version", "string", "资源版本", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func versionList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := chi.URLParam(r, "resource")
		list, err := service.New().Resource().VersionList(r.Context(), resource)
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
				Kind:       v.Kind,
				Format:     v.Format,
				Desc:       v.Desc,
				Content:    v.Data,
				Timestamp: response.Timestamp{
					CreatedTS: v.CreatedAt.Unix(),
					UpdatedTS: v.UpdatedAt.Unix(),
				},
			})
		}
		ctr.OK(w, result)
	}
}

func versionInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := chi.URLParam(r, "resource")
		version := chi.URLParam(r, "version")
		info, err := service.New().Resource().VersionFind(r.Context(), resource, version)
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
			Content:    info.Data,
			Timestamp: response.Timestamp{
				CreatedTS: info.CreatedAt.Unix(),
				UpdatedTS: info.UpdatedAt.Unix(),
			},
		})
	}
}

func versionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := chi.URLParam(r, "resource")
		version := chi.URLParam(r, "version")
		if err := service.New().Resource().VersionDelete(r.Context(), resource, version); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
