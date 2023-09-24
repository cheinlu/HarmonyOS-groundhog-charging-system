package v2_saas

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta   `path:"user/list" tags:"用户管理" method:"get" summary:"用户列表"`
	Username string `json:"username"`
	model.PageReq
}

type UserListRes struct {
	g.Meta     `mime:"application/json" `
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Balance    float64   `json:"balance"  ` //
	RoleName   string    `json:"roleName"`
	CreateTime time.Time `json:"createTime"`
}

type UserAddReq struct {
	g.Meta   `path:"user/add" tags:"用户管理" method:"POST" summary:"用户添加"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserAddRes struct {
	g.Meta `mime:"application/json" `
}

type UserDelReq struct {
	g.Meta `path:"user/delete" tags:"用户管理" method:"DELETE" summary:"用删除"`
	Id     int `json:"id"`
}

type UserDelRes struct {
	g.Meta `mime:"application/json" `
}

type UserUpdateReq struct {
	g.Meta   `path:"user/update" tags:"用户管理" method:"post" summary:"用户修改"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserUpdateRes struct {
	g.Meta `mime:"application/json" `
}
