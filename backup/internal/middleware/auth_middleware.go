package middleware

import (
	"errors"
	"login-demo/internal/logic"
	"login-demo/internal/packed/jwt"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 检测登录 token 是否合法
func Auth(r *ghttp.Request) {
	// 从 HTTP 请求的 Header 中获取 Token
	token, err := getToken(r)
	if err != nil {
		handleError(r, 401, err, "Token 验证失败")
	}
	// 验证 Token 是否有效
	if isValidToken := jwt.Valid(r.Context(), token); !isValidToken {
		handleError(r, 401, err, "Token 验证失败")
	}

	// 解析 token 中 user 的信息存入 Ctx 中
	if err := parseTokenAndSetContext(r, token); err != nil {
		handleError(r, 401, err, "Token 验证失败")
		return
	}
	r.Middleware.Next()
}

// handleError 处理错误并设置响应
func handleError(r *ghttp.Request, code int, err error, detail string) {
	errMsg := detail
	if err != nil {
		e := gerror.WrapCode(gcode.New(code, err.Error(), detail), err)
		r.SetError(e)
	} else {
		e := gerror.NewCode(gcode.New(code, errMsg, detail))
		r.SetError(e)
	}
}

// parseTokenAndSetContext 解析 token 并将用户信息设置到上下文中
func parseTokenAndSetContext(r *ghttp.Request, token string) error {
	claims, err := jwt.Parse(r.Context(), token)
	if err != nil {
		return err
	}
	username := claims.Username
	logic.User.SetUserContext(username, r.SetCtxVar)
	return nil
}

// GetToken 从 HTTP 请求的 Header 中获取 Token
func getToken(r *ghttp.Request) (string, error) {
	// 从 HTTP 请求的 Header 中获取 Authorization Header
	header := r.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("token 验证失败，未找到 Authorization Header")
	}
	// 将 Authorization Header 拆分为 Token 类型和 Token 值
	headerList := strings.Split(header, " ")
	// 如果拆分后的数组长度不为 2，说明 Authorization Header 格式不正确，返回 false
	if len(headerList) != 2 {
		return "", errors.New("token 验证失败， Authorization Header 格式不正确")
	}
	// 获取 Token 类型
	t := headerList[0]
	// 获取 Token 值
	token := headerList[1]
	// 如果 Token 类型不为 "Bearer"，说明 Token 格式不正确，返回 false
	if t != "Bearer" {
		return "", errors.New("token 验证失败，token 类型不为 Bearer")
	}
	// 如果 Token 值为空，说明 Token 格式不正确，返回 false
	if token == "" {
		return "", errors.New("token 验证失败，token 为空")
	}
	// 返回 Token 值和 true，表示获取 Token 成功
	return token, nil
}
