package controller

import (
	"context"
	v1 "login-demo/api/v1"
	"login-demo/internal/logic"
)

// var Login CLogin = CLogin{}

type Login struct {
}

// 登录
func (c Login) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	role, tokenString, err := logic.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	res = &v1.LoginRes{
		Token: tokenString,
		Role:  role,
	}
	return
}

func (C Login) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	newToken, err := logic.MyJwt.GenerateToken(ctx, logic.Ctx.GetHTTPContextParams(ctx).Username)
	res = &v1.RefreshRes{
		Token: newToken,
	}
	return
}
