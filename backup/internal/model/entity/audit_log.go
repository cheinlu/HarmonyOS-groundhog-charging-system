// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLog is the golang structure for table audit_log.
type AuditLog struct {
	Id            int         `json:"id"            ` //
	TenantId      int         `json:"tenantId"      ` //
	OperationType string      `json:"operationType" ` //
	TableName     string      `json:"tableName"     ` //
	OldValue      string      `json:"oldValue"      ` //
	NewValue      string      `json:"newValue"      ` //
	OperateSql    string      `json:"operateSql"    ` //
	ChangedDate   *gtime.Time `json:"changedDate"   ` //
	ChangedUser   string      `json:"changedUser"   ` //
}
