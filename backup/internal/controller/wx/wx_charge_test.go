package controller_wx_test

import (
	"context"
	v2 "login-demo/api/wx"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/wx"
	"login-demo/internal/model"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestWxStationList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	gtest.C(t, func(t *gtest.T) {
		page, err := controller.WxChargeController{}.StationList(context.Background(), &v2.WxStationListReq{
			// Name: "admin",
			// Address: "admin",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 3)
		list, ok := page.List.([]*v2.WxStationListRes)
		t.AssertEQ(ok, true)
		t.Assert(list[0].TenantName, "土拨鼠充电公司")
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.WxChargeController{}.StationList(context.Background(), &v2.WxStationListReq{
			// Name: "admin",
			Address: "宝安",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 2)
	})

	// 查询
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.WxChargeController{}.StationList(context.Background(), &v2.WxStationListReq{
			// Name: "admin",
			Address: "黑龙江",
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 0)
	})
}

func TestWxPileList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.WxChargeController{}.PileList(ctx, &v2.WxPileListReq{
			StationId: 1,
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 4)
	})
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.WxChargeController{}.PileList(ctx, &v2.WxPileListReq{
			StationId: 2,
		})
		t.AssertNil(err)
		t.Assert(page.TotalCount, 1)
	})
}

func TestWxStartCharge(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		_, err := controller.WxChargeController{}.StartCharge(userCtx, &v2.WXStartChargeReq{
			PileId:      1,
			ChargeHours: 4,
		})
		t.AssertNil(err)
	})
	// 登录用户不存在
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.WxChargeController{}.StartCharge(ctx, &v2.WXStartChargeReq{
			PileId:      1,
			ChargeHours: 2,
		})
		t.AssertNE(err, nil)
	})
	// 充电桩状态异常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		_, err := controller.WxChargeController{}.StartCharge(userCtx, &v2.WXStartChargeReq{
			PileId:      4,
			ChargeHours: 2,
		})
		t.AssertNE(err, nil)
	})
}

func TestWxStopCharge(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		_, err := controller.WxChargeController{}.StopCharge(userCtx, &v2.WXStopChargeReq{
			OrderId: 1,
		})
		t.AssertNil(err)
	})
	// 登录用户不存在
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.WxChargeController{}.StopCharge(ctx, &v2.WXStopChargeReq{
			OrderId: 1,
		})
		t.AssertNE(err, nil)
	})
	// 充电桩状态异常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		_, err := controller.WxChargeController{}.StopCharge(userCtx, &v2.WXStopChargeReq{
			OrderId: 2,
		})
		t.AssertNE(err, nil)
	})
}

func TestMyChargeOrders(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		pageRes, err := controller.WxChargeController{}.MyChargeOrders(userCtx, &v2.WXChargeListReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 2)
		list, ok := pageRes.List.([]v2.WXChargeListRes)
		t.Assert(ok, true)
		t.Assert(list[0].OrderCode, "202305111249-8745")
	})
	// 查询 guest 的订单
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "guest",
		})
		pageRes, err := controller.WxChargeController{}.MyChargeOrders(userCtx, &v2.WXChargeListReq{})
		t.Assert(err, nil)
		t.Assert(pageRes.TotalCount, 1)
	})
	// 登录用户不存在
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.WxChargeController{}.MyChargeOrders(ctx, &v2.WXChargeListReq{})
		t.AssertNE(err, nil)
	})
}

func TestAboutMe(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		res, err := controller.WxChargeController{}.AboutMe(userCtx, &v2.WXMeInfoReq{})
		t.AssertNil(err)
		t.Assert(res.Username, "admin")
	})
	// 查询 guest 的订单
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "guest",
		})
		res, err := controller.WxChargeController{}.AboutMe(userCtx, &v2.WXMeInfoReq{})
		t.AssertNil(err)
		t.Assert(res.Username, "guest")
	})
	// 登录用户不存在
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.WxChargeController{}.AboutMe(ctx, &v2.WXMeInfoReq{})
		t.AssertNE(err, nil)
	})
}

func TestPriceList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 正常
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "admin",
		})
		pageRes, err := controller.WxChargeController{}.PriceList(userCtx, &v2.WXPriceListReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 6)
		list, ok := pageRes.List.([]v2.WXPriceListRes)
		t.Assert(ok, true)
		t.Assert(list[0].Price, 0.7)
		t.Assert(list[0].TenantId, 1)
	})
	// 查询 guest 的订单
	gtest.C(t, func(t *gtest.T) {
		userCtx := context.WithValue(ctx, consts.UserContextKey, model.UserContext{
			Username: "guest",
		})
		pageRes, err := controller.WxChargeController{}.PriceList(userCtx, &v2.WXPriceListReq{})
		t.AssertNil(err)
		t.Assert(pageRes.TotalCount, 6)
		list, ok := pageRes.List.([]v2.WXPriceListRes)
		t.Assert(ok, true)
		t.Assert(list[0].Price, 0.7)
		t.Assert(list[0].TenantId, 1)
	})
}
