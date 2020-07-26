/**
 * Created by zc on 2020/6/10.
 */
package api

import (
	"github.com/go-chi/chi"
	"luban/global"
	"luban/pkg/ctr"
	"luban/pkg/errs"
	"luban/service"
	"net/http"
)

// JwtAuth returns a handler to verify user token
func JwtAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(global.HeaderTokenKey)
		var claims ctr.JwtClaims
		if err := ctr.JwtParse(&claims, token, global.Cfg().Server.Secret); err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		if err := claims.Valid(); err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		ctx := ctr.ContextWithUser(r.Context(), &claims.User)
		users, err := service.New().User().All(ctx)
		if err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		var exist bool
		for _, user := range users {
			if user.Code == claims.User.Code {
				exist = true
				break
			}
		}
		if !exist {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(errs.New("user not exist")))
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// SpaceAuth returns a handler to verify space value
func SpaceAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		space := chi.URLParam(r, "space")
		if space == "" {
			ctr.BadRequest(w, errs.ErrInvalidSpace)
			return
		}
		ctx := ctr.ContextWithSpace(r.Context(), space)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// ConfigAuth returns a handler to verify config value
func ConfigAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		config := chi.URLParam(r, "config")
		if config == "" {
			ctr.BadRequest(w, errs.ErrInvalidConfig)
			return
		}
		ctx := ctr.ContextWithConfig(r.Context(), config)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
