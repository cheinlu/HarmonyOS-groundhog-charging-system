// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWallet is the golang structure of table user_wallet for DAO operations like Where/Data.
type UserWallet struct {
	g.Meta   `orm:"table:user_wallet, do:true"`
	UserId   interface{} //
	TenantId interface{} //
	Balance  interface{} //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
