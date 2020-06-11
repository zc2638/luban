/**
 * Created by zc on 2020/6/9.
 */
package auth

import (
	"net/http"
	"stone/global"
	"stone/pkg/api"
	"stone/pkg/api/store"
	"stone/pkg/ctr"
	"stone/pkg/errs"
	"stone/service"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.RegisterParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		// check if the user name is duplicate
		user, ok := service.New().User().FindByEmail(params.Email)
		if ok {
			ctr.BadRequest(w, errs.New("Duplicate email"))
			return
		}
		user = &store.User{
			Username: params.Username,
			Email:    params.Email,
			Pwd:      params.Password,
		}
		if err := service.New().User().Create(user); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Str(w, "ok")
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params api.LoginParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		user, err := service.New().User().FindByNameAndPwd(params.Username, params.Password)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		userInfo := ctr.JwtUserInfo{
			UID:      user.UID,
			Username: user.Username,
			Email:    user.Email,
		}
		token, err := ctr.JwtCreate(ctr.JwtClaims{User: userInfo}, global.Cfg().Serve.Secret)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, api.LoginResult{
			Username: user.Username,
			Email:    user.Email,
			Token:    token,
		})
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := ctr.ContextUserFrom(r.Context())
		if !ok {
			ctr.Unauthorized(w, errs.ErrUnauthorized)
			return
		}
		ctr.OK(w, user)
	}
}