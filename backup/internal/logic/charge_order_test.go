package logic_test

import (
	"context"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestOrderList(t *testing.T) {
	createTableFromSql()
	defer dropAllTableWithDb()
	gtest.C(t, func(t *gtest.T) {
		orders, count, err := logic.ChargeOrder.ChargeOrderList(context.Background(), model.ChargeOrderQuery{})
		t.AssertNil(err)
		t.Assert(count, 3)
		t.Assert(orders[0].OrderCode, "202305111249-8745")
	})
}
