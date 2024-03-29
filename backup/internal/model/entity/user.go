// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       ` //
	TenantId int         `json:"tenantId" ` //
	RoleId   int         `json:"roleId"   ` //
	Passport string      `json:"passport" ` //
	Password string      `json:"password" ` //
	Nickname string      `json:"nickname" ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
}
