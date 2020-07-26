/**
 * Created by zc on 2020/6/6.
 */
package web

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"path/filepath"
	"strings"
)

func New() http.Handler {
	mux := chi.NewMux()
	mux.Mount("/", fileSystem())
	return mux
}

func fileSystem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
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