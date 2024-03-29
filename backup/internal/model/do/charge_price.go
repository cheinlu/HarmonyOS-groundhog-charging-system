// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChargePrice is the golang structure of table charge_price for DAO operations like Where/Data.
type ChargePrice struct {
	g.Meta    `orm:"table:charge_price, do:true"`
	Id        interface{} //
	TenantId  interface{} //
	StartHour interface{} //
	EndHour   interface{} //
	Price     interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
}
