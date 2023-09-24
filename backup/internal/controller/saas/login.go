package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/packed/jwt"

	"github.com/mojocn/base64Captcha"
)

// var Login CLogin = CLogin{}

type Login struct {
}

var captcha = base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)

// 登录
func (c Login) Login(ctx context.Context, req *v2.LoginReq) (res *v2.LoginRes, err error) {
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

// func (C Login) GetCaptcha(ctx context.Context, req *v2.GetCaptchaReq) (res *v2.GetCaptchaRes, err error) {

// 	id, b64s, err := captcha.Generate()
// 	if err != nil {
// 		// err = gerror.NewCode(gcode.New(code, message, detail))
// 	}
// 	res = &v2.GetCaptchaRes{
// 		CaptchaId:   id,
// 		CaptchaBody: b64s,
// 	}
// 	return
// }

// func (C Login) ValidCaptcha(ctx context.Context, req *v2.ValidCaptchaReq) (res *v2.ValidCaptchaRes, err error) {
// 	res = &v2.ValidCaptchaRes{}
// 	return
// }
