package sql_test

import (
	"login-demo/internal/packed/sql"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestAddWhere(t *testing.T) {
	// 没有 where 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select * from user", nil)
		parser.AddWhere("tenant_id", 1)
		t.Assert(parser.GetSQL(), "select * from user where tenant_id=?")
		t.Assert(len(parser.GetArgs()), 1)
		t.Assert(parser.GetArgs()[0], 1)
	})

	// 没有 where，有 as
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select * from user as a", nil)
		parser.AddWhere("a.tenant_id", 1)
		t.Assert(parser.GetSQL(), "select * from user as a where a.tenant_id=?")
		t.Assert(len(parser.GetArgs()), 1)
		t.Assert(parser.GetArgs()[0], 1)
	})

	// 有 where 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select * from user where i=?", []interface{}{100})
		parser.AddWhere("tenant_id", 1)
		t.Assert(parser.GetSQL(), "select * from user where tenant_id=? and i=?")
		t.Assert(len(parser.GetArgs()), 2)
		t.Assert(parser.GetArgs()[0], 1)
		t.Assert(parser.GetArgs()[1], 100)
	})

	// 有 where 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select * from user where i=?", []interface{}{100})
		parser.AddWhere("tenant_id", 1)
		t.Assert(parser.GetSQL(), "select * from user where tenant_id=? and i=?")
		t.Assert(len(parser.GetArgs()), 2)
		t.Assert(parser.GetArgs()[0], 1)
		t.Assert(parser.GetArgs()[1], 100)
	})

	// 有 order by 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select * from user where i=? order by id desc", []interface{}{100})
		parser.AddWhere("tenant_id", 1)
		t.Assert(parser.GetSQL(), "select * from user where tenant_id=? and i=? order by id desc")
		t.Assert(len(parser.GetArgs()), 2)
		t.Assert(parser.GetArgs()[0], 1)
		t.Assert(parser.GetArgs()[1], 100)
	})

	// 有 join 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select user.*,role.name from user join role where user.role_id = role.id and user.id=? order by id desc", []interface{}{100})
		parser.AddWhere("user.tenant_id", 1)
		t.Assert(parser.GetSQL(), "select user.*,role.name from user join role where user.tenant_id=? and user.role_id = role.id and user.id=? order by id desc")
		t.Assert(len(parser.GetArgs()), 2)
		t.Assert(parser.GetArgs()[0], 1)
		t.Assert(parser.GetArgs()[1], 100)
	})

	// 有 left join 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("select user.*,role.name from user left join role where user.role_id = role.id order by id desc", nil)
		parser.AddWhere("user.tenant_id", 1)
		t.Assert(parser.GetSQL(), "select user.*,role.name from user left join role where user.tenant_id=? and user.role_id = role.id order by id desc")
		t.Assert(len(parser.GetArgs()), 1)
		t.Assert(parser.GetArgs()[0], 1)
	})

	// 有 join on 和 order 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("Select user.*,role.name From user Join role On user.role_id = role.id Order by id desc", nil)
		parser.AddWhere("user.tenant_id", 1)
		t.Assert(parser.GetSQL(), "Select user.*,role.name From user Join role On user.role_id = role.id where user.tenant_id=? Order by id desc")
		t.Assert(len(parser.GetArgs()), 1)
		t.Assert(parser.GetArgs()[0], 1)
	})

	// 有 join on 和 group by 的情况
	gtest.C(t, func(t *gtest.T) {
		parser := sql.NewSQLParser("Select user.*,role.name From user Join role On user.role_id = role.id group by id", nil)
		parser.AddWhere("user.tenant_id", 1)
		t.Assert(parser.GetSQL(), "Select user.*,role.name From user Join role On user.role_id = role.id where user.tenant_id=? group by id")
		t.Assert(len(parser.GetArgs()), 1)
		t.Assert(parser.GetArgs()[0], 1)
	})
}
