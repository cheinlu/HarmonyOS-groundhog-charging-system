// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure of table tenant for DAO operations like Where/Data.
type Tenant struct {
	g.Meta       `orm:"table:tenant, do:true"`
	Id           interface{} //
	Name         interface{} //
	Domain       interface{} //
	ContactName  interface{} //
	ContactEmail interface{} //
	ContactPhone interface{} //
	Status       interface{} //
	CreateAt     *gtime.Time //
	UpdateAt     *gtime.Time //
}
