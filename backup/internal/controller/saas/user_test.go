package controller_saas_test

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestUserList(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 查看所有 tenant = 1 的 station
	gtest.C(t, func(t *gtest.T) {
		page, err := controller.UserController{}.List(tenantCtx, &v2.UserListReq{})
		t.Assert(page.TotalCount, 3)
		t.Assert(err, nil)
		userList, ok := page.List.([]v2.UserListRes)
		t.Assert(ok, true)
		t.Assert(userList[0].RoleName, "超级管理员")
	})
}

func TestUserAdd(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Add(tenantCtx, &v2.UserAddReq{
			Username: "admin123",
			Password: "123456",
		})
		t.AssertNil(err)
	})
	// 2. 不存在 tenantId，添加失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Add(context.Background(), &v2.UserAddReq{
			Username: "admin123",
			Password: "123456",
		})
		t.AssertNE(err, nil)
	})
	// 3. 用户名已经存在，添加失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Add(context.Background(), &v2.UserAddReq{
			Username: "admin",
			Password: "123456",
		})
		t.AssertNE(err, nil)
	})
}

func TestUserDel(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Del(tenantCtx, &v2.UserDelReq{
			Id: 1,
		})
		t.AssertNil(err)
	})
	// 2. id 不存在，失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Del(tenantCtx, &v2.UserDelReq{
			Id: 10,
		})
		t.AssertNE(err, nil)
	})
}

func TestUserUpdate(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	tenantCtx := context.WithValue(context.Background(), consts.TenantIDKey, 1)

	// 1. 存在 tenantId，添加成功
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Update(tenantCtx, &v2.UserUpdateReq{
			Id:       1,
			Username: "admin@qq.com",
		})
		t.AssertNil(err)
	})
	// 2. id 不存在，失败
	gtest.C(t, func(t *gtest.T) {
		_, err := controller.UserController{}.Update(tenantCtx, &v2.UserUpdateReq{
			Id:       10,
			Username: "admin@qq.com",
		})
		t.AssertNE(err, nil)
	})
}
