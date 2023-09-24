package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestPriceList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		page, err := controller.ChargePriceController{}.List(ctx, &v2.ChargePriceListReq{})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 3)
		list, ok := page.List.([]v2.ChargePriceListRes)
		t.Assert(ok, true)
		t.Assert(list[0].StartHour, 0)
		t.Assert(list[0].TenantId, 1)
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 2)
		page, err := controller.ChargePriceController{}.List(ctx, &v2.ChargePriceListReq{})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 2)
		list, ok := page.List.([]v2.ChargePriceListRes)
		t.Assert(ok, true)
		t.Assert(list[0].StartHour, 0)
		t.Assert(list[0].TenantId, 2)
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 3)
		page, err := controller.ChargePriceController{}.List(ctx, &v2.ChargePriceListReq{})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 1)
		list, ok := page.List.([]v2.ChargePriceListRes)
		t.Assert(ok, true)
		t.Assert(list[0].StartHour, 0)
		t.Assert(list[0].TenantId, 3)
	})
}

func TestPriceAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(ctx, consts.TenantIDKey, 1)
		_, err := controller.ChargePriceController{}.Add(tenantCtx, &v2.ChargePriceAddReq{
			StartHour: 1,
			EndHour:   10,
			Price:     5.1,
		})
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.ChargePriceController{}.Add(context.Background(), &v2.ChargePriceAddReq{
			StartHour: 1,
			EndHour:   10,
			Price:     5.1,
		})
		t.AssertNE(err, nil)
	})
}

func TestPriceDel(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	//
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(ctx, consts.TenantIDKey, 1)
		_, err := controller.ChargePriceController{}.Del(tenantCtx, &v2.ChargePriceDelReq{
			Id: 1,
		})
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.ChargePriceController{}.Del(ctx, &v2.ChargePriceDelReq{
			Id: 10,
		})
		t.AssertNE(err, nil)
	})
}

func TestPriceUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	//
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(ctx, consts.TenantIDKey, 1)
		_, err := controller.ChargePriceController{}.Update(tenantCtx, &v2.ChargePriceUpdateReq{
			Id:        1,
			StartHour: 10,
			EndHour:   20,
			Price:     1,
		})
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.ChargePriceController{}.Update(ctx, &v2.ChargePriceUpdateReq{
			Id:        10,
			StartHour: 10,
			EndHour:   20,
			Price:     1,
		})
		t.AssertNE(err, nil)
	})
}
