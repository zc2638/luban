/**
 * Created by zc on 2020/6/9.
 */
package auth

import (
	"context"
	"github.com/pkgms/go/ctr"
	"github.com/zc2638/gotool/utilx"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"luban/global"
	"luban/pkg/api/request"
	"luban/pkg/api/response"
	"luban/pkg/compile"
	"luban/pkg/store"
	"luban/pkg/wrapper"
	"luban/service"
	"net/http"
)

// Route handle auth routing related
func Route(doc *swagger.API) {
	const tag = "auth"
	doc.Tags = append(doc.Tags, swagger.Tag{
		Name:        tag,
		Description: "用户认证",
	})
	doc.AddEndpoint(
		endpoint.New(
			http.MethodPost, "/auth/register",
			endpoint.Handler(register()),
			endpoint.Summary("用户注册"),
			endpoint.Body(request.RegisterParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodPost, "/auth/login",
			endpoint.Handler(login()),
			endpoint.Summary("用户登陆"),
			endpoint.Body(request.LoginParams{}, "", true),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
		endpoint.New(
			http.MethodGet, "/auth/info",
			endpoint.Handler(info()),
			endpoint.Summary("用户信息"),
			endpoint.ResponseSuccess(),
			endpoint.Tags(tag),
		),
	)
}

func register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.RegisterParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		if !compile.Username().MatchString(params.Username) {
			ctr.BadRequest(w, compile.UsernameError)
			return
		}
		user := &store.User{
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

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.LoginParams
		if err := wrapper.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		user, err := service.New().User().FindByNameAndPwd(context.Background(), params.Username, params.Password)
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		userInfo := wrapper.JwtUserInfo{
			UserID:   user.UserID,
			Username: user.Username,
			Pwd:      user.Pwd,
		}
		token, err := wrapper.JwtCreate(wrapper.JwtClaims{User: userInfo}, global.Cfg().Server.Secret)
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

func info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := wrapper.ContextUserFrom(r.Context())
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
