package controller_saas_test

import (
	"context"
	"testing"

	v2 "login-demo/api/saas"
	"login-demo/internal/consts"
	controller "login-demo/internal/controller/saas"
	"login-demo/internal/logic"
	_ "login-demo/internal/packed/driver"
	"login-demo/internal/packed/jwt"

	"github.com/gogf/gf/v2/test/gtest"

	_ "login-demo/internal/packed"
)

func TestLogin(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()
	ctx := context.WithValue(context.Background(), consts.TenantIDKey, consts.DefaultTenantValue)

	// 1. 登录成功
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.Login{}.Login(ctx, &v2.LoginReq{
			Username: "admin",
			Password: "admin",
		})
		t.AssertNil(err)
		t.AssertNE(res.Token, "")
	})

	// 2. 登录失败，未知账户和密码
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.Login{}.Login(ctx, &v2.LoginReq{
			Username: "coding",
			Password: "coding123",
		})
		t.Assert(res, nil)
		t.AssertNE(err, nil)
	})

	// 3. 登录失败，租户密码不匹配
	gtest.C(t, func(t *gtest.T) {
		res, err := controller.Login{}.Login(context.Background(), &v2.LoginReq{
			Username: "admin",
			Password: "123456",
		})
		t.Assert(res, nil)
		t.AssertNE(err, nil)
	})
}

func TestRefresh(t *testing.T) {
	CreateTableFromSql()
	defer DropAllTableWithDb()

	res, err := controller.Login{}.Login(context.Background(), &v2.LoginReq{
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		t.Error(err)
	}
	token := res.Token
	// 正确的
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		userClaim, err := jwt.Parse(ctx, token)
		t.Assert(err, nil)
		setCtx := func(key, value interface{}) {
			ctx = context.WithValue(ctx, key, value)
		}
		logic.User.SetUserContext(userClaim.Username, setCtx)
		res, err := controller.Login{}.Refresh(ctx, &v2.RefreshReq{})
		t.AssertNil(err)
		t.AssertNE(res.Token, "")
	})

	// 错误的
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		res, err := controller.Login{}.Refresh(ctx, &v2.RefreshReq{})
		t.AssertNE(err, nil)
		t.AssertNil(res)
	})

}
