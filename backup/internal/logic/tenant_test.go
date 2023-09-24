package logic_test

import (
	"context"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/model/do"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestSetCurrentTenantId(t *testing.T) {
	ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
	// 1. 存在 tenantId
	gtest.C(t, func(t *gtest.T) {
		station := do.Station{}
		logic.SetCurrentTenantId(ctx, &station.TenantId)
		t.Assert(station.TenantId, 1)
	})
	// 2. 不存在 tenantId，异常
	gtest.C(t, func(t *gtest.T) {
		station := do.Station{}
		err := logic.SetCurrentTenantId(context.Background(), &station.TenantId)
		t.AssertNE(err, nil)
	})
}
