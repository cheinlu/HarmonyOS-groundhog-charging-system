package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"login-demo/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestAuditLogList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	gtest.C(t, func(t *gtest.T) {
		pageRes, err := controller.AuditLogController{}.List(ctx, &v2.AuditLogListReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 2)
		auditLogs, ok := pageRes.List.([]entity.AuditLog)
		t.Assert(ok, true)
		t.Assert(len(auditLogs), 2)
		t.Assert(auditLogs[0].OperationType, "add")
	})

	gtest.C(t, func(t *gtest.T) {
		pageRes, err := controller.AuditLogController{}.List(ctx, &v2.AuditLogListReq{
			OperationType: "add",
			ChangedUser:   "admin",
		})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 1)
		auditLogs, ok := pageRes.List.([]entity.AuditLog)
		t.Assert(ok, true)
		t.Assert(len(auditLogs), 1)
		t.Assert(auditLogs[0].OperationType, "add")
	})
}
