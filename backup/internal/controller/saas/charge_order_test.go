package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestOrderList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		page, err := controller.ChargeOrderController{}.List(ctx, &v2.ChargeOrderListReq{})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 3)
		list, ok := page.List.([]v2.ChargeOrderListRes)
		t.Assert(ok, true)
		t.Assert(list[0].OrderCode, "202305111249-8745")
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.ChargeOrderController{}.List(context.Background(), &v2.ChargeOrderListReq{
			OrderCode: "202305111249",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 1)
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 2)
		page, err := controller.ChargeOrderController{}.List(ctx, &v2.ChargeOrderListReq{
			// Name: "admin",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 0)
	})
}

func TestOrderUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		_, err := controller.ChargeOrderController{}.Update(ctx, &v2.ChargeOrderUpdateReq{
			Id:    1,
			Price: 10,
			State: 1,
		})
		t.AssertNil(err)
	})

	// id不存在
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		_, err := controller.ChargeOrderController{}.Update(ctx, &v2.ChargeOrderUpdateReq{
			Id:    10,
			Price: 10,
			State: 1,
		})
		t.AssertNE(err, nil)
	})

	// 租户不对查询
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 2)
		_, err := controller.ChargeOrderController{}.Update(ctx, &v2.ChargeOrderUpdateReq{
			Id:    10,
			Price: 10,
			State: 1,
		})
		t.AssertNE(err, nil)
	})

}
