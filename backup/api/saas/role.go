package v2_saas

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"role/list" tags:"角色管理" method:"get" summary:"角色列表"`
	Name   string `json:"name"`
	model.PageReq
}

type RoleListRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BindUsers int       `json:"bindUsers"`    //已经绑定的用户数
	CreateAt  time.Time `json:"createAt"    ` //
	UpdateAt  time.Time `json:"updateAt"    ` //
}

type RoleAddReq struct {
	g.Meta `path:"role/add" tags:"角色管理" method:"POST" summary:"添加角色"`
	Name   string `json:"name"`
}

type RoleAddRes struct {
	g.Meta `mime:"application/json" `
}

type RoleDelReq struct {
	g.Meta `path:"role/delete" tags:"角色管理" method:"DELETE" summary:"删除角色"`
	Id     int `json:"id"`
}

type RoleDelRes struct {
	g.Meta `mime:"application/json" `
}

type RoleUpdateReq struct {
	g.Meta `path:"role/update" tags:"角色管理" method:"post" summary:"修改角色"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

type RoleUpdateRes struct {
	g.Meta `mime:"application/json" `
}
