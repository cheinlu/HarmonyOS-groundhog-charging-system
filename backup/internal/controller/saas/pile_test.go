package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestPileList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 1. 带租户id查询
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.PileController{}.List(ctx, &v2.PileListReq{})
		t.Assert(err, nil)
		t.Assert(page.TotalCount, 6)
	})
	// 2. 没有的租户id，查询全部
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.PileController{}.List(context.Background(), &v2.PileListReq{})
		t.Assert(err, nil)
		t.Assert(page.TotalCount, 6)
	})
}

func TestPileAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 1. 带租户id查询
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
		_, err := controller.PileController{}.Add(tenantCtx, &v2.PileAddReq{
			Code:      "288",
			StationId: 12,
		})
		t.Assert(err, nil)
	})

	// 2. 不带租户id查询
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
		_, err := controller.PileController{}.Add(tenantCtx, &v2.PileAddReq{
			Code:      "288",
			StationId: 12,
		})
		t.AssertNE(err, nil)
	})
}

func TestPileUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 1. 带租户id查询
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
		_, err := controller.PileController{}.Update(tenantCtx, &v2.PileUpdateReq{
			Id:        1,
			Code:      "288",
			StationId: 12,
			State:     1,
		})
		t.Assert(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 10)
		_, err := controller.PileController{}.Update(tenantCtx, &v2.PileUpdateReq{
			Id:        10,
			Code:      "288",
			StationId: 12,
			State:     1,
		})
		t.AssertNE(err, nil)
	})
}

func TestPileDel(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 1. 带租户id查询
	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
		_, err := controller.PileController{}.Del(tenantCtx, &v2.PileDelReq{
			Id: 1,
		})
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
		_, err := controller.PileController{}.Del(tenantCtx, &v2.PileDelReq{
			Id: 10,
		})
		t.Assert(err, "删除失败，未找到原数据，可能已被其他人删除")
	})
}
