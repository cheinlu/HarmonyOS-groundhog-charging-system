package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/api/user/login" tags:"login" method:"POST" summary:"登录"`
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
	g.Meta `path:"/api/token/refresh" tags:"refresh" method:"POST" summary:"刷新凭证"`
}

type RefreshRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Token  string `json:"token"`
	Role   string `json:"role" summary:"user or admin"`
}
