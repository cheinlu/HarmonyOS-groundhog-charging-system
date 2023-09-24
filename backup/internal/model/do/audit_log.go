// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLog is the golang structure of table audit_log for DAO operations like Where/Data.
type AuditLog struct {
	g.Meta        `orm:"table:audit_log, do:true"`
	Id            interface{} //
	TenantId      interface{} //
	OperationType interface{} //
	TableName     interface{} //
	OldValue      interface{} //
	NewValue      interface{} //
	OperateSql    interface{} //
	ChangedDate   *gtime.Time //
	ChangedUser   interface{} //
}
