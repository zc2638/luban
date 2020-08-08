/**
 * Created by zc on 2020/6/9.
 */
package auth

import (
	"context"
	"github.com/zc2638/gotool/utilx"
	"luban/global"
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/service"
	"net/http"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.RegisterParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Username().MatchString(params.Username) {
			ctr.BadRequest(w, compile.UsernameError)
			return
		}
		user := &data.User{
			Username: params.Username,
			Pwd:      params.Password,
			Salt:     utilx.RandomStr(6),
		}
		if err := service.New().User().Create(context.Background(), user); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.LoginParams
		if err := ctr.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		user, err := service.New().User().FindByNameAndPwd(context.Background(), params.Username, params.Password)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		userInfo := ctr.JwtUserInfo{
			UserID:   user.UserID,
			Username: user.Username,
			Pwd:      user.Pwd,
		}
		token, err := ctr.JwtCreate(ctr.JwtClaims{User: userInfo}, global.Cfg().Server.Secret)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, response.LoginResult{
			Username: user.Username,
			Token:    token,
			Host:     global.Cfg().Server.Host,
		})
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := ctr.ContextUserFrom(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, response.UserInfoResult{
			Username: user.Username,
			Host:     global.Cfg().Server.Host,
		})
	}
}
