package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestStationList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 查看所有 tenant = 1 的 station
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.StationController{}.List(ctx, &v2.StationListReq{
			// Name: "admin",
			// Address: "admin",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 2)
		list, ok := page.List.([]*v2.StationListRes)
		t.Assert(ok, true)
		t.Assert(list[0].TenantName, "土拨鼠充电公司")
	})

	// 查询 Address 的 station
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.StationController{}.List(ctx, &v2.StationListReq{
			// Name: "admin",
			Address: "宝安",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 2)
	})

	// 查询一个不存在的 station
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.StationController{}.List(ctx, &v2.StationListReq{
			// Name: "admin",
			Address: "黑龙江",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 0)
	})
}

func TestStationAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Add(tenantCtx, &v2.StationAddReq{
			Name:     "充电站A",
			Address:  "北京",
			ImageUrl: "xx.png",
		})
		t.AssertNil(err)
	})
	// 2. 不存在 tenantId，添加失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Add(context.Background(), &v2.StationAddReq{
			Name:    "充电站A",
			Address: "北京",
		})
		t.AssertNE(err, nil)
	})
}

func TestStationUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Update(tenantCtx, &v2.StationUpdateReq{
			Id:       1,
			Name:     "充电站A",
			Address:  "北京",
			ImageUrl: "xx.png",
		})
		t.AssertNil(err)
	})
	// 2. 不存在 tenantId，添加失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Update(context.Background(), &v2.StationUpdateReq{
			Id:      10,
			Name:    "充电站A",
			Address: "北京",
		})
		t.AssertNE(err, nil)
	})
}

func TestStationDel(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Del(tenantCtx, &v2.StationDelReq{
			Id: 1,
		})
		t.AssertNil(err)
	})
	// 2. 不存在 tenantId，添加失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.StationController{}.Del(context.Background(), &v2.StationDelReq{
			Id: 10,
		})
		t.AssertNE(err, nil)
	})
}
