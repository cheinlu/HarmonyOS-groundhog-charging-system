package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"login-demo/internal/logic"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

// todo: 改造saas后，该单元测试存在问题
func TestPermissionList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// tenant == 1
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 1)
		res, err := controller.PermissionController{}.List(ctx, &v2.PermissionListReq{})
		t.AssertNil(err)
		t.Assert(len(res.Data), 11)
	})
	// tenant == 2
	gtest.C(t, func(t *gtest.T) {
		ctx = context.WithValue(ctx, consts.TenantIDKey, 2)
		res, err := controller.PermissionController{}.List(ctx, &v2.PermissionListReq{})
		t.AssertNil(err)
		t.Assert(len(res.Data), 10)
	})
}

func TestPermissionBind(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	// 有 tenantId，正常
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.PermissionController{}.Bind(ctx, &v2.BindPermissionReq{
			RoleId:         2,
			PermissionList: []string{"ChargeStationView"},
		})
		t.AssertNil(err)
	})
	// 没有 tenantId，报错
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.PermissionController{}.Bind(context.Background(), &v2.BindPermissionReq{
			RoleId:         2,
			PermissionList: []string{"ChargeStationView"},
		})
		t.AssertNE(err, nil)
	})
}

func TestCurrentPermissions(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)
	logic.User.SetUserContext("admin", func(i1, i2 interface{}) {
		ctx = context.WithValue(ctx, i1, i2)
	})
	// 有 tenantId，正常
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.PermissionController{}.CurrentPermissions(ctx, &v2.CurrentPermissionReq{})
		t.AssertNil(err)
		t.AssertNE(len(res.PermissionList), 0)
	})
}
