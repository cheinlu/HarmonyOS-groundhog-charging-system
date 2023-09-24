package driver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"login-demo/internal/consts"
	"login-demo/internal/model"
	"strings"
	"time"

	"github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
)

// MyDriver is a custom database driver, which is used for testing only.
// For simplifying the unit testing case purpose, MyDriver struct inherits the mysql driver
// gdb.Driver and overwrites its functions DoQuery and DoExec.
// So if there's any sql execution, it goes through MyDriver.DoQuery/MyDriver.DoExec firstly
// and then gdb.Driver.DoQuery/gdb.Driver.DoExec.
// You can call it sql "HOOK" or "HiJack" as your will.
type MyDriver struct {
	*sqlite.Driver
}

var (
	// customDriverName is my driver name, which is used for registering.
	customDriverName = "MyDriver"
)

func init() {
	// It here registers my custom driver in package initialization function "init".
	// You can later use this type in the database configuration.
	if err := gdb.Register(customDriverName, &MyDriver{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *MyDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &MyDriver{
		&sqlite.Driver{
			Core: core,
		},
	}, nil
}

// DoCommit commits current sql and arguments to underlying sql driver.
func (d *MyDriver) DoCommit(ctx context.Context, in gdb.DoCommitInput) (out gdb.DoCommitOutput, err error) {
	// glog.Debugf(ctx, "===============type=%s,sql=%s,db=%v,tran=%t,args=%v\n", in.Type, in.Sql, in.Db, in.IsTransaction, in.Args)
	tenantId, ok := ctx.Value(consts.TenantIDKey).(int)
	userCtx, userOk := ctx.Value(consts.UserContextKey).(model.UserContext)

	if ok && userOk {
		operationType := strings.ToLower(strings.Split(in.Sql, " ")[0])
		if operationType == "update" || operationType == "insert" || operationType == "delete" || operationType == "replace" {
			argStrs := make([]string, len(in.Args))
			for i, v := range in.Args {
				argStrs[i] = fmt.Sprint(v)
			}
			argStr := strings.Join(argStrs, ", ")
			// tode oldValue
			// var oldValueStr string
			// if operationType == "update" {
			// 	orginsql := fmt.Sprintf("select * from %s where %s", strings.Split(in.Sql, " ")[1], strings.Split(strings.ToLower(in.Sql), "where")[1])
			// 	argCount := strings.Count(orginsql, "?")
			// 	rows, _ := in.Link.QueryContext(ctx, orginsql, in.Args[len(in.Args)-argCount:])
			// 	oldValueStr, _ = RowsToJSONArray(rows)
			// } else if operationType == "delete" {
			// 	orginsql := fmt.Sprintf("select * from %s where %s", strings.Split(in.Sql, " ")[2], strings.Split(strings.ToLower(in.Sql), "where")[1])
			// 	// argCount := strings.Count(orginsql, "?")
			// 	rows, _ := in.Link.QueryContext(ctx, orginsql, in.Args)
			// 	oldValueStr, _ = RowsToJSONArray(rows)
			// }
			_, err = in.Link.ExecContext(ctx,
				"INSERT INTO audit_log(tenant_id, operation_type, table_name, old_value, new_value, operate_sql, changed_date, changed_user) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
				tenantId,
				operationType,
				strings.Split(in.Sql, " ")[1],
				"",
				argStr,
				in.Sql,
				time.Now(),
				userCtx.Username,
			)
			glog.Debugf(ctx, "err=%v\n", err)
		}
	}

	out, err = d.Core.DoCommit(ctx, in)
	return
}

func RowsToJSONArray(rows *sql.Rows) (jsonStr string, err error) {
	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}

		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		}

		result := make(map[string]interface{})
		for i, column := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				val = string(b)
			}
			result[column] = val
		}

		results = append(results, result)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr = string(jsonData)
	return
}
