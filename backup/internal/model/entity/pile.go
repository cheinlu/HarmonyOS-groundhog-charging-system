// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Pile is the golang structure for table pile.
type Pile struct {
	Id        int         `json:"id"        ` //
	Code      string      `json:"code"      ` //
	StationId int         `json:"stationId" ` //
	State     int         `json:"state"     ` //
	TenantId  int         `json:"tenantId"  ` //
	CreateAt  *gtime.Time `json:"createAt"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  ` //
}
