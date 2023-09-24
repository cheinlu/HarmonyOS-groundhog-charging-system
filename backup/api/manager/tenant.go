package v2_manager

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type TenantSelectListReq struct {
	g.Meta `path:"/tenant/selects" tags:"租户管理" method:"get" summary:"租户列表（简略）"`
	model.PageReq
}

type TenantSelectListRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []TenantSelectList
}

type TenantSelectList struct {
	Id     int    `json:"id"           ` //
	Name   string `json:"name"         ` //
	Domain string `json:"domain"       ` //
}

type TenantRegisterReq struct {
	g.Meta        `path:"/tenant/register" tags:"租户管理" method:"post" summary:"租户注册"`
	Name          string `json:"name"         `        //
	Domain        string `json:"domain"       `        //
	AdminUsername string `json:"adminUsername"       ` //
	AdminPassword string `json:"adminPassword"       ` //
	ValidCode     string `json:"ValidCode"       `     //
}

type TenantRegisterRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type TenantListReq struct {
	g.Meta `path:"/tenant/list" tags:"租户管理" method:"get" summary:"租户列表"`
	model.PageReq
}

type TenantListRes struct {
	g.Meta       `mime:"application/json" example:"string"`
	Id           int       `json:"id"           ` //
	Name         string    `json:"name"         ` //
	Domain       string    `json:"domain"       ` //
	ContactName  string    `json:"contactName"  ` //
	ContactEmail string    `json:"contactEmail" ` //
	ContactPhone string    `json:"contactPhone" ` //
	Status       string    `json:"status"       ` //
	CreateAt     time.Time `json:"createAt"     ` //
	UpdateAt     time.Time `json:"updateAt"     ` //
}

type TenantAddReq struct {
	g.Meta       `path:"/tenant/add" tags:"租户管理" method:"POST" summary:"添加租户"`
	Name         string `json:"name"         ` //
	Domain       string `json:"domain"       ` //
	ContactName  string `json:"contactName"  ` //
	ContactEmail string `json:"contactEmail" ` //
	ContactPhone string `json:"contactPhone" ` //
	Status       string `json:"status"       ` //
}

type TenantAddRes struct {
	g.Meta `mime:"application/json" `
}

type TenantDelReq struct {
	g.Meta `path:"/tenant/delete" tags:"租户管理" method:"DELETE" summary:"删除租户"`
	Id     int `json:"id"`
}

type TenantDelRes struct {
	g.Meta `mime:"application/json" `
}

type TenantUpdateReq struct {
	g.Meta       `path:"/tenant/update" tags:"租户管理" method:"post" summary:"修改租户"`
	Id           int    `json:"id"`
	Name         string `json:"name"         ` //
	Domain       string `json:"domain"       ` //
	ContactName  string `json:"contactName"  ` //
	ContactEmail string `json:"contactEmail" ` //
	ContactPhone string `json:"contactPhone" ` //
	Status       string `json:"status"       ` //
}

type TenantUpdateRes struct {
	g.Meta `mime:"application/json" `
}
