/**
 * Created by zc on 2020/6/10.
 */
package api

import (
	"luban/global"
	"luban/pkg/ctr"
	"luban/pkg/errs"
	"net/http"
)

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
		next.ServeHTTP(w, r.WithContext(
			ctr.ContextWithUser(r.Context(), &claims.User),
		))
	}
	return http.HandlerFunc(fn)
}
