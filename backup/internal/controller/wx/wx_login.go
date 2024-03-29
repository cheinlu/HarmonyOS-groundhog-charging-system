package controller_wx

import (
	"context"
	"database/sql"
	v2 "login-demo/api/wx"
	"login-demo/internal/logic"
	"login-demo/internal/packed/jwt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type WxLoginController struct {
}

func (WxLoginController) Login(ctx context.Context, req *v2.WxLoginReq) (res *v2.WxLoginRes, err error) {
	code := req.Code
	data, err := logic.WxLogin.WechatLogin(ctx, code)
	if err != nil {
		return nil, err
	}
	// 将 openid 和 session_key 存储到数据库或缓存中等待后续使用
	user, err := logic.WxUser.GetUserByOpenID(ctx, data.OpenID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = gerror.NewCode(gcode.New(404, "微信登录失败", "用户未注册"))
			return
		}
		return
	}
	token, err := jwt.GenerateToken(ctx, user.Passport)
	if err != nil {
		return nil, err
	}
	res = &v2.WxLoginRes{
		Token: token,
	}
	return
}

// 注册接口
func (WxLoginController) Register(ctx context.Context, req *v2.WxRegisterReq) (res *v2.WxRegisterRes, err error) {
	// 验证验证码
	// err = logic.WxRegister.VerifyCode(ctx, req.Phone, req.VerifyCode)
	// if err != nil {
	// 	err = gerror.NewCode(gcode.New(1, "验证码验证失败", err.Error()))
	// 	return
	// }

	// 获得微信用户的 openId
	data, err := logic.WxLogin.WechatLogin(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 查询用户是否已经存在
	_, err = logic.WxUser.GetUserByOpenID(ctx, data.OpenID)
	if err == nil {
		return nil, gerror.NewCode(gcode.New(1, "用户已经存在，无需注册", "用户已经存在，无需注册"))
	}

	// 注册新用户
	token, err := logic.WxRegister.Register(ctx, req.Nickname, req.AvatarUrl, data.OpenID, req.Gender)
	if err != nil {
		return nil, err
	}

	res = &v2.WxRegisterRes{
		Token: token,
	}

	return res, nil
}

// 获取手机号，这个接口无法使用，因为个人小程序不允许获取用户手机号
func (WxLoginController) GetPhone(ctx context.Context, req *v2.WxGetPhoneReq) (res *v2.WxGetPhoneRes, err error) {
	res = &v2.WxGetPhoneRes{}
	phone, err := logic.WxLogin.WechatPhone(ctx, req.Code, req.EncryptedData, req.Iv)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "获取微信手机号失败", err.Error()))
		return
	}
	res.Phone = phone.PhoneInfo.PhoneNumber
	return
}

// 注册验证码
func (WxLoginController) SendValidCode(ctx context.Context, req *v2.WxValidCodeReq) (res *v2.WxValidCodeRes, err error) {
	err = logic.WxLogin.SendValidCode(ctx, req.Phone)
	return
}
