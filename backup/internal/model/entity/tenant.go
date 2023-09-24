// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure for table tenant.
type Tenant struct {
	Id           int         `json:"id"           ` //
	Name         string      `json:"name"         ` //
	Domain       string      `json:"domain"       ` //
	ContactName  string      `json:"contactName"  ` //
	ContactEmail string      `json:"contactEmail" ` //
	ContactPhone string      `json:"contactPhone" ` //
	Status       string      `json:"status"       ` //
	CreateAt     *gtime.Time `json:"createAt"     ` //
	UpdateAt     *gtime.Time `json:"updateAt"     ` //
}
