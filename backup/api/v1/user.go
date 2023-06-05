package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta   `path:"/api/user/list" tags:"login" method:"get" summary:"用户列表"`
	Username string `json:"username"`
	model.PageReq
}

type UserListRes struct {
	g.Meta     `mime:"application/json" `
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Balance    float64   `json:"balance"  ` //
	CreateTime time.Time `json:"createTime"`
}

type UserAddReq struct {
	g.Meta   `path:"/api/user/add" tags:"userAdd" method:"POST" summary:"登录"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserAddRes struct {
	g.Meta `mime:"application/json" `
}

type UserDelReq struct {
	g.Meta `path:"/api/user/delete" tags:"userDelete" method:"DELETE" summary:"删除"`
	Id     int `json:"id"`
}

type UserDelRes struct {
	g.Meta `mime:"application/json" `
}

type UserUpdateReq struct {
	g.Meta   `path:"/api/user/update" tags:"userDelete" method:"post" summary:"修改"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserUpdateRes struct {
	g.Meta `mime:"application/json" `
}
