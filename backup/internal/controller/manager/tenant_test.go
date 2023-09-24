package controller_manager_test

import (
	"context"
	v2 "login-demo/api/manager"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/manager"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestTenantList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	gtest.C(t, func(t *gtest.T) {
		pageRes, err := controller.TenantController{}.List(ctx, &v2.TenantListReq{})
		t.Assert(err, nil)
		tenantlist, ok := pageRes.List.([]*v2.TenantListRes)
		t.Assert(ok, true)
		t.Assert(len(tenantlist), 2)
		t.Assert(tenantlist[0].Status, "active")
	})
}

func TestTenantSelectList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	gtest.C(t, func(t *gtest.T) {
		selects, err := controller.TenantController{}.SelectList(ctx, &v2.TenantSelectListReq{})
		t.Assert(2, len(selects.Data))
		t.AssertNil(err)
		t.Assert(selects.Data[0].Name, "土拨鼠充电公司")
	})
}

func TestTenantRegister(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.Background() //注册时，没有租户
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.TenantController{}.Register(ctx, &v2.TenantRegisterReq{
			Name:          "北京新东方有限公司",
			Domain:        "xdf.com",
			AdminUsername: "xdf",
			AdminPassword: "123456",
			ValidCode:     "666666",
		})
		t.Log(err)
		t.Assert(err, nil)
	})
}

func TestTenantAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.TenantController{}.Add(ctx, &v2.TenantAddReq{
			Name:         "北京新东方有限公司",
			Domain:       "xdf.com",
			ContactName:  "张三",
			ContactEmail: "x@qq.com",
			ContactPhone: "13422334422",
			Status:       "active",
		})
		t.Log(err)
		t.Assert(err, nil)
	})
}

func TestTenantUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.TenantController{}.Update(ctx, &v2.TenantUpdateReq{
			Id:           1,
			Name:         "北京新东方有限公司",
			Domain:       "xdf.com",
			ContactName:  "张三",
			ContactEmail: "x@qq.com",
			ContactPhone: "13422334422",
			Status:       "active",
		})
		t.Log(err)
		t.Assert(err, nil)
	})
}

func TestTenantDel(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.TenantController{}.Del(ctx, &v2.TenantDelReq{
			Id: 1,
		})
		t.Log(err)
		t.Assert(err, nil)
	})
}
