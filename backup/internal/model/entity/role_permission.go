// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermission is the golang structure for table role_permission.
type RolePermission struct {
	RoleId      int         `json:"roleId"      ` //
	TenantId    int         `json:"tenantId"    ` //
	Permissions string      `json:"permissions" ` //
	CreateAt    *gtime.Time `json:"createAt"    ` //
	UpdateAt    *gtime.Time `json:"updateAt"    ` //
}
