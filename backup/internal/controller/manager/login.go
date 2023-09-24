package controller_manager

import (
	"context"
	v2 "login-demo/api/manager"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/packed/jwt"
)

// var Login CLogin = CLogin{}

type Login struct {
}

// 登录
func (c Login) Login(ctx context.Context, req *v2.LoginReq) (res *v2.LoginRes, err error) {
	ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
	role, tokenString, err := logic.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	res = &v2.LoginRes{
		Token: tokenString,
		Role:  role,
	}
	return
}

func (C Login) Refresh(ctx context.Context, req *v2.RefreshReq) (res *v2.RefreshRes, err error) {
	user, err := logic.User.GetCurrentUser(ctx)
	if err != nil {
		return
	}
	newToken, err := jwt.GenerateToken(ctx, user.Passport)
	res = &v2.RefreshRes{
		Token: newToken,
	}
	return
}
