package internal

import (
	"context"
	"fmt"
	"login-demo/internal/consts"
	"login-demo/internal/packed/sql"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

var TenantHook = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		// 忽略的表名
		ignoredTables := []string{"tenant"}
		// 判断当前表是否需要执行 hook
		v := reflect.ValueOf(in.Model).Elem()
		field := v.FieldByName("tablesInit")
		if !field.IsValid() {
			err = gerror.NewCode(gcode.New(1, "未找到 model(%s) 表名", ""), "")
			return
		}
		table := strings.Trim(field.String(), "`")
		// 从请求的上下文中获取租户ID
		tenantID, ok := ctx.Value(consts.TenantIDKey).(int)
		if !contains(ignoredTables, table) && ok {
			// 构造租户查询条件
			sqlParser := sql.NewSQLParser(in.Sql, in.Args)
			tenantCondition := "tenant_id"
			// 如果是联表查询，表名.tenant_id
			if strings.Contains(strings.ToLower(in.Sql), "join") {
				tenantCondition = fmt.Sprintf("%s.tenant_id", table)
			}
			sqlParser.AddWhere(tenantCondition, tenantID)
			glog.Debugf(ctx, "租户 hook 变更执行sql，变更前sql: %s, 参数%v\n", in.Sql, in.Args)
			in.Sql = sqlParser.GetSQL()
			in.Args = sqlParser.GetArgs()
			glog.Debugf(ctx, "租户 hook 变更执行sql，变更后sql: %s, 参数%v\n", in.Sql, in.Args)
		}
		result, err = in.Next(ctx)
		return
	},
}

// 判断字符串切片中是否包含指定字符串
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
