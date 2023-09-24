package v2_saas

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionListReq struct {
	g.Meta `path:"permission/list" tags:"权限" method:"GET" summary:"权限列表"`
}

type PermissionListRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []PermissionList `json:"data"`
}

type PermissionList struct {
	Model          string       `json:"model"`
	PermissionList []Permission `json:"permissionList"`
}

type Permission struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type BindPermissionReq struct {
	g.Meta         `path:"permission/bind" tags:"权限" method:"POST" summary:"角色绑定权限"`
	RoleId         int      `json:"roleId"`
	PermissionList []string `json:"permissionList"`
}

type BindPermissionRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type CurrentPermissionReq struct {
	g.Meta `path:"permission/current" tags:"权限" method:"GET" summary:"当前权限"`
}

type CurrentPermissionRes struct {
	g.Meta         `mime:"application/json" example:"string"`
	PermissionList []string `json:"permissionList"`
}
