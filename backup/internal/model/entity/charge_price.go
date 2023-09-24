// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChargePrice is the golang structure for table charge_price.
type ChargePrice struct {
	Id        int         `json:"id"        ` //
	TenantId  int         `json:"tenantId"  ` //
	StartHour int         `json:"startHour" ` //
	EndHour   int         `json:"endHour"   ` //
	Price     float64     `json:"price"     ` //
	CreateAt  *gtime.Time `json:"createAt"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  ` //
}
