// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WxUser is the golang structure of table wx_user for DAO operations like Where/Data.
type WxUser struct {
	g.Meta    `orm:"table:wx_user, do:true"`
	UserId    interface{} //
	TenantId  interface{} //
	OpenId    interface{} //
	PhoneNo   interface{} //
	AvatarUrl interface{} //
	Nickname  interface{} //
	Gender    interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
}
