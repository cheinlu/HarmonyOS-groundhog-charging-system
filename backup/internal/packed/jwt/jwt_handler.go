package jwt

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt"
)

var (
	JwtSecret   = []byte("123456")
	ExpiresTime = time.Minute * 1 // 默认1分钟
)

func init() {
	val, _ := g.Cfg().Get(context.Background(), "tokenAliveTime")
	tokenAlive := val.Int()
	ExpiresTime = time.Minute * time.Duration(tokenAlive)
}

type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成 jwt 格式 token
func GenerateToken(ctx context.Context, username string) (token string, err error) {
	tokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpiresTime).Unix(),
		},
	})
	token, err = tokenHeader.SignedString(JwtSecret)
	return
}

// Valid 验证 token 是否合法
func Valid(ctx context.Context, token string) (valid bool) {
	var claims *UserClaims = &UserClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	valid = (err == nil) && tkn.Valid
	return
}

// 解析 token 内容
func Parse(ctx context.Context, token string) (claims *UserClaims, err error) {
	claims = &UserClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	return
}
