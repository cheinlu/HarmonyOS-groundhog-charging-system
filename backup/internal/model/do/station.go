// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Station is the golang structure of table station for DAO operations like Where/Data.
type Station struct {
	g.Meta     `orm:"table:station, do:true"`
	Id         interface{} //
	TenantId   interface{} //
	Name       interface{} //
	Address    interface{} //
	ImageUrl   interface{} //
	Coordinate interface{} //
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
}
