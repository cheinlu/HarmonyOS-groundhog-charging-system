// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChargeOrder is the golang structure for table charge_order.
type ChargeOrder struct {
	Id        int         `json:"id"        ` //
	TenantId  int         `json:"tenantId"  ` //
	OrderCode string      `json:"orderCode" ` //
	UserId    int         `json:"userId"    ` //
	StationId int         `json:"stationId" ` //
	PileId    int         `json:"pileId"    ` //
	StartAt   *gtime.Time `json:"startAt"   ` //
	StopAt    *gtime.Time `json:"stopAt"    ` //
	State     int         `json:"state"     ` //
	Price     float64     `json:"price"     ` //
	CreateAt  *gtime.Time `json:"createAt"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  ` //
}
