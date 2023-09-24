package logic_test

import (
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestAuditList(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		auditLogs, count, err := logic.AuditLogLogic{}.List(ctx, model.PageReq{}, entity.AuditLog{})
		t.AssertNil(err)
		t.Assert(count, 2)
		t.Assert(auditLogs[0].OperationType, "add")
	})
}
