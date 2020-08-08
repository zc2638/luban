/**
 * Created by zc on 2020/8/8.
**/
package route

import (
	"github.com/go-chi/chi"
	"luban/handler/api"
	"luban/handler/web"
	"net/http"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/user/login/", http.StatusFound)
	})
	mux.Mount("/v1", api.New())
	mux.Mount("/web", web.New())
	mux.Handle("/favicon.png", web.FaviconImage())
	return mux
}