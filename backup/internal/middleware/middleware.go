package middleware

import (
	"login-demo/internal/logic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 解析 Jwt 的信息存入 Ctx 中
func ParseJwtToCtx(r *ghttp.Request) {
	_, username := logic.User.Parse(r.GetCtx(), r)
	logic.Ctx.SetHTTPContextParams(r, username)
	r.Middleware.Next()
}

// 检测登录 token 是否合法
func Auth(r *ghttp.Request) {
	valid := logic.User.IsSignedIn(r.GetCtx(), r)
	if !valid {
		r.SetError(gerror.NewCode(gcode.New(50008, "请重新登录", "非法令牌或还未登录")))
		return
	}
	r.Middleware.Next()
}
