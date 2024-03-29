// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRecord is the golang structure of table pay_record for DAO operations like Where/Data.
type PayRecord struct {
	g.Meta   `orm:"table:pay_record, do:true"`
	Id       interface{} //
	TenantId interface{} //
	PayCode  interface{} //
	UserId   interface{} //
	State    interface{} //
	Price    interface{} //
	PayAt    *gtime.Time //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
