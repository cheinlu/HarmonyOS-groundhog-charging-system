// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRecord is the golang structure for table pay_record.
type PayRecord struct {
	Id       int         `json:"id"       ` //
	TenantId int         `json:"tenantId" ` //
	PayCode  string      `json:"payCode"  ` //
	UserId   int         `json:"userId"   ` //
	State    int         `json:"state"    ` //
	Price    float64     `json:"price"    ` //
	PayAt    *gtime.Time `json:"payAt"    ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
}
