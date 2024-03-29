// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Pile is the golang structure of table pile for DAO operations like Where/Data.
type Pile struct {
	g.Meta    `orm:"table:pile, do:true"`
	Id        interface{} //
	Code      interface{} //
	StationId interface{} //
	State     interface{} //
	TenantId  interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
}
