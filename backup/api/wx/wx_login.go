package v2_wx

import "github.com/gogf/gf/v2/frame/g"

type WxLoginReq struct {
	g.Meta `path:"/wx-api/login" tags:"微信小程序鉴权接口" method:"POST" summary:"微信登录"`
	Code   string `json:"code"`
}

type WxLoginRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
}

type WxGetPhoneReq struct {
	g.Meta        `path:"/wx-api/getphone" tags:"微信小程序鉴权接口" method:"POST" summary:"获得手机号"`
	Code          string `json:"code"`
	Iv            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WxGetPhoneRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Phone  string `json:"phone"`
}

type WxValidCodeReq struct {
	g.Meta `path:"/wx-api/validcode" tags:"微信小程序鉴权接口" method:"GET" summary:"发送验证码"`
	Phone  string `json:"phone"`
}

type WxValidCodeRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type WxRegisterReq struct {
	g.Meta `path:"/wx-api/register" tags:"微信小程序鉴权接口" method:"POST" summary:"微信注册"`
	Code   string `json:"code"`
	// Phone      string `json:"phone"`
	// VerifyCode string `json:"verifyCode"`
	Gender    string `json:"gender"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type WxRegisterRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
}
