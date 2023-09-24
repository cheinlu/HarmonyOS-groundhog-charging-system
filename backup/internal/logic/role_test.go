package logic_test

import (
	"context"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestRoleList(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		roles, count, err := logic.Role.RoleList(ctx, "管理员", model.PageReq{
			PageNo:   1,
			PageSize: 5,
		})
		t.Assert(len(roles), 2)
		t.Assert(len(roles), count)
		t.Assert(roles[0].TenantId, 1)
		t.AssertNil(err)
	})
}
