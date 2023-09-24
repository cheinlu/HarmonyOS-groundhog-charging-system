package logic_test

import (
	"context"
	"login-demo/internal/consts"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestGetCurrentUser(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.UserContextKey, model.UserContext{
		Username: "test",
	})
	gtest.C(t, func(t *gtest.T) {
		user, err := logic.User.GetCurrentUser(ctx)
		t.AssertNil(err)
		t.Assert(user.Passport, "test")
	})
}
