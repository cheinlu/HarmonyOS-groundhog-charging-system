// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChargeOrder is the golang structure of table charge_order for DAO operations like Where/Data.
type ChargeOrder struct {
	g.Meta    `orm:"table:charge_order, do:true"`
	Id        interface{} //
	TenantId  interface{} //
	OrderCode interface{} //
	UserId    interface{} //
	StationId interface{} //
	PileId    interface{} //
	StartAt   *gtime.Time //
	StopAt    *gtime.Time //
	State     interface{} //
	Price     interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
}
