// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"login-demo/internal/dao/internal"
)

// internalTenantDao is internal type for wrapping internal DAO implements.
type internalTenantDao = *internal.TenantDao

// tenantDao is the data access object for table tenant.
// You can define custom methods on it to extend its functionality as you wish.
type tenantDao struct {
	internalTenantDao
}

var (
	// Tenant is globally public accessible object for table tenant operations.
	Tenant = tenantDao{
		internal.NewTenantDao(),
	}
)

// Fill with you ideas below.
