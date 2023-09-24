package logic_test

import (
	"context"
	"login-demo/internal/logic"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestWxRegister(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()
	ctx := context.Background()
	gtest.C(t, func(t *gtest.T) {
		token, err := logic.WxRegister.Register(ctx, "张三", "", "4444", "男")
		t.AssertNil(err)
		t.AssertNE(token, "")
	})
}
