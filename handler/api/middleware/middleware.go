/**
 * Created by zc on 2020/6/10.
 */
package middleware

import (
	"context"
	"net/http"
	"stone/global"
	"stone/pkg/ctr"
	"stone/pkg/errs"
)

func JwtAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		claims, err := ctr.JwtParse(r.Header.Get(global.HeaderTokenKey), global.Cfg().Serve.Secret)
		if err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		if err := claims.Valid(); err != nil {
			ctr.Unauthorized(w, errs.ErrInvalidToken.With(err))
			return
		}
		next.ServeHTTP(w, r.WithContext(
			ctr.WithUser(context.Background(), &claims.User),
		))
	}
	return http.HandlerFunc(fn)
}
