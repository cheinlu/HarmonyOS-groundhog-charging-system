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

func TestRoleList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	// 1. 带租户id查询
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.RoleController{}.List(ctx, &v2.RoleListReq{})
		t.Assert(page.TotalCount, 3)
		list, ok := page.List.([]*v2.RoleListRes)
		t.Assert(ok, true)
		t.Assert(list[0].BindUsers, 1)
		t.AssertNil(err)
	})
	// 2. 带租户id不存在的查询
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.Background(), consts.TenantIDKey, 2)
		page, err := controller.RoleController{}.List(ctx, &v2.RoleListReq{})
		t.Assert(page.TotalCount, 0)
		t.AssertNil(err)
	})
	// 3. 查询不到 name
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.RoleController{}.List(ctx, &v2.RoleListReq{
			Name: "终极顾客",
		})
		t.Assert(page.TotalCount, 0)
		t.AssertNil(err)
	})
}

// 角色添加
func TestRoleAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
	// 正常添加
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Add(ctx, &v2.RoleAddReq{
			Name: "测试角色",
		})
		t.AssertNil(err)
	})

	// 添加不带 name，失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Add(ctx, &v2.RoleAddReq{})
		t.AssertNE(err, nil)
	})

	// 添加不带 tenantId，失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Add(ctx, &v2.RoleAddReq{})
		t.AssertNE(err, nil)
	})
}

// 删除角色
func TestRoleDel(t *testing.T) {
	// 删除为1的
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
	logic.User.SetUserContext("admin", func(i1, i2 interface{}) {
		ctx = context.WithValue(ctx, i1, i2)
	})
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Del(ctx, &v2.RoleDelReq{
			Id: 1,
		})
		t.AssertNil(err)
	})

	// 删除为10的，没有
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Del(ctx, &v2.RoleDelReq{
			Id: 10,
		})
		t.AssertNil(err)
	})

	// 删除没有传 id，报错
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Del(ctx, &v2.RoleDelReq{})
		t.AssertNE(err, nil)
	})
}

func TestRoleUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, 1)
	logic.User.SetUserContext("admin", func(i1, i2 interface{}) {
		ctx = context.WithValue(ctx, i1, i2)
	})
	// 1. 修改为1的
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Update(ctx, &v2.RoleUpdateReq{
			Id:   1,
			Name: "测试管理员",
		})
		t.AssertNil(err)
	})

	// 2. 修改为10的，没有
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Update(ctx, &v2.RoleUpdateReq{
			Id:   10,
			Name: "测试管理员",
		})
		t.AssertNil(err)
	})

	// 3. 修改没有传id，报错
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.RoleController{}.Update(ctx, &v2.RoleUpdateReq{
			Name: "测试管理员",
		})
		t.AssertNE(err, nil)
	})
}
