// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id       int         `json:"id"       ` //
	TenantId int         `json:"tenantId" ` //
	Name     string      `json:"name"     ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
}
