/**
 * Created by zc on 2020/8/8.
**/
package route

import (
	"github.com/go-chi/chi"
	"github.com/zc2638/drone-control/handler"
	"luban/handler/api"
	"luban/handler/web"
	"net/http"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web", http.StatusFound)
	})
	mux.Mount("/v1", api.New())
	mux.Mount("/web", web.New())
	mux.Handle("/favicon.png", web.FaviconImage())
	mux.Mount("/rpc/v2", handler.RPC())
	return mux
}