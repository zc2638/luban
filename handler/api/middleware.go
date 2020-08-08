/**
 * Created by zc on 2020/6/10.
 */
package api

import (
	"github.com/go-chi/chi"
	"luban/global"
	"luban/pkg/compile"
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
		user, err := service.New().User().FindByUserID(ctx, claims.User.UserID)
		if err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(errs.New("user not exist")))
			return
		}
		if user.Pwd != claims.User.Pwd {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(errs.New("user password is out")))
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
		if space == "" || !compile.Name().MatchString(space) {
			ctr.BadRequest(w, errs.ErrInvalidSpace)
			return
		}
		ctx := ctr.ContextWithSpace(r.Context(), space)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// ResourceAuth returns a handler to verify config value
func ResourceAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resource := chi.URLParam(r, "resource")
		if resource == "" || !compile.Name().MatchString(resource) {
			ctr.BadRequest(w, errs.ErrInvalidResource)
			return
		}
		ctx := ctr.ContextWithResource(r.Context(), resource)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
