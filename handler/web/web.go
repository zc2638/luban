/**
 * Created by zc on 2020/6/6.
 */
package web

import (
	"github.com/go-chi/chi"
	"net/http"
	"path/filepath"
	"strings"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Mount("/", fileSystem())
	//mux.Handle("/", http.FileServer(http.Dir("public/dist")))
	//mux.Handle("/static", http.FileServer(http.Dir("public/dist/static")))
	//mux.Handle("/icons", http.FileServer(http.Dir("public/dist/icons")))
	return mux
}

func fileSystem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := strings.Trim(r.URL.Path, "/web")
		list := strings.Split(url, "/")
		current := make([]string, 0, len(list))
		for _, v := range list {
			if v == "" {
				continue
			}
			current = append(current, v)
		}
		fp := "public/dist"
		cLen := len(current)
		if cLen > 0 {
			p := current[len(current) - 1]
			if filepath.Ext(p) != "" {
				if cLen > 1 && current[len(current)-2] == "static" {
					fp = filepath.Join(fp, "static")
				}
				fp = filepath.Join(fp, p)
				http.ServeFile(w, r, fp)
				return
			}
		}
		http.ServeFile(w, r, fp)
	}
}

func FaviconImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/dist/favicon.png")
	}
}