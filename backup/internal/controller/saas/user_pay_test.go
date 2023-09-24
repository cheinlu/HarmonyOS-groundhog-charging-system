package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	"testing"

	controller "login-demo/internal/controller/saas"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestUserPayList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)

	// 查看所有 tenant = 1 的 user 支付订单
	gtest.C(t, func(t *gtest.T) {
		pageRes, err := controller.UserPayController{}.List(tenantCtx, &v2.UserPayRecordReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 4)
		list, ok := pageRes.List.([]v2.UserPayRecordRes)
		t.Assert(ok, true)
		t.Assert(list[0].PayCode, "202305111249-8745")
	})
	// 查看所有 tenant = 2 的 user 支付订单
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 2)
		pageRes, err := controller.UserPayController{}.List(tenantCtx, &v2.UserPayRecordReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 0)
	})
}

func TestUserPayAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)

	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserPayController{}.Pay(tenantCtx, &v2.UserPayReq{
			UserId: 1,
			Price:  10,
		})
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserPayController{}.Pay(tenantCtx, &v2.UserPayReq{
			UserId: 10,
			Price:  10,
		})
		t.Assert(err, "用户不存在")
	})
}
