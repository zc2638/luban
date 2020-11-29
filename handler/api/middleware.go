/**
 * Created by zc on 2020/6/10.
 */
package api

import (
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"luban/global"
	"luban/pkg/wrapper"
	"luban/pkg/errs"
	"luban/service"
	"net/http"
)

// JwtAuth returns a handler to verify user token
func JwtAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(global.HeaderTokenKey)
		if token == "" {
			next.ServeHTTP(w, r.WithContext(
				wrapper.ContextWithUser(r.Context(), &wrapper.JwtUserInfo{})))
			return
		}
		var claims wrapper.JwtClaims
		if err := wrapper.JwtParse(&claims, token, global.Cfg().Server.Secret); err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		if err := claims.Valid(); err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		ctx := wrapper.ContextWithUser(r.Context(), &claims.User)
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
		if space == "" {
			ctr.BadRequest(w, errs.ErrInvalidSpace)
			return
		}
		ctx := wrapper.ContextWithSpace(r.Context(), space)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// ResourceAuth returns a handler to verify resource value
func ResourceAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resource := chi.URLParam(r, "resource")
		if resource == "" {
			ctr.BadRequest(w, errs.ErrInvalidResource)
			return
		}
		ctx := wrapper.ContextWithResource(r.Context(), resource)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// PipelineAuth returns a handler to verify pipeline value
func PipelineAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pipeline := chi.URLParam(r, "pipeline")
		if pipeline == "" {
			ctr.BadRequest(w, errs.ErrInvalidPipeline)
			return
		}
		ctx := wrapper.ContextWithPipeline(r.Context(), pipeline)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// TaskAuth returns a handler to verify task value
func TaskAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		task := chi.URLParam(r, "task")
		if task == "" {
			ctr.BadRequest(w, errs.ErrInvalidTask)
			return
		}
		ctx := wrapper.ContextWithTask(r.Context(), task)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
