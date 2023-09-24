package v2_saas

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" tags:"SaaS：登录鉴权" method:"POST" summary:"登录"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
	// user or admin
	Role string `json:"role"`
}

type RefreshReq struct {
	g.Meta `path:"/token/refresh" tags:"SaaS：登录鉴权" method:"POST" summary:"刷新凭证"`
}

type RefreshRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
	Role   string `json:"role" summary:"user or admin"`
}

// type GetCaptchaReq struct {
// 	g.Meta `path:"/captcha/get" tags:"SaaS：登录鉴权" method:"GET" summary:"图片验证码"`
// }

// type GetCaptchaRes struct {
// 	g.Meta      `mime:"application/json" example:"string"`
// 	CaptchaId   string `json:"captcha_id"`
// 	CaptchaBody string `json:"captcha_body" summary:"图片内容，需要 base64 解码"`
// }

// type ValidCaptchaReq struct {
// 	g.Meta      `path:"/captcha/valid" tags:"SaaS：登录鉴权" method:"POST" summary:"验证码效验"`
// 	CaptchaId   string `json:"captcha_id"`
// 	CaptchaCode string `json:"captcha_code"`
// }

// type ValidCaptchaRes struct {
// 	g.Meta `mime:"application/json" example:"string"`
// }
