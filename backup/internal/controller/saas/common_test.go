package controller_saas_test

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "login-demo/internal/packed/driver"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	db         gdb.DB
	dbPrefix   gdb.DB
	dbInvalid  gdb.DB
	configNode gdb.ConfigNode
	dbDir      = gfile.Temp("sqlite")
	ctx        = gctx.New()

	// Error
	ErrorSave = gerror.NewCode(gcode.CodeNotSupported, `Save operation is not supported by sqlite driver`)
)

const (
	TableNamePrefix = "gf_"
	DBGroupTest     = "test"
	DBGroupPrefix   = "prefix"
	DBGroupInvalid  = "invalid"
	InitSql         = "../../../manifest/db/db.sql"
)

func init() {
	if err := gfile.Mkdir(dbDir); err != nil {
		gtest.Error(err)
	}

	fmt.Println("init sqlite db dir: ", dbDir)

	dbFilePath := gfile.Join(dbDir, "test.db")
	configNode = gdb.ConfigNode{
		Type:    "sqlite",
		Link:    fmt.Sprintf(`MyDriver::@file(%s)`, dbFilePath),
		Charset: "utf8",
	}
	nodePrefix := configNode
	nodePrefix.Prefix = TableNamePrefix

	nodeInvalid := configNode

	gdb.AddConfigNode(DBGroupTest, configNode)
	gdb.AddConfigNode(DBGroupPrefix, nodePrefix)
	gdb.AddConfigNode(DBGroupInvalid, nodeInvalid)
	gdb.AddConfigNode(gdb.DefaultGroupName, configNode)

	// Default db.
	if r, err := gdb.NewByGroup(); err != nil {
		gtest.Error(err)
	} else {
		db = r
	}

	// Prefix db.
	if r, err := gdb.NewByGroup(DBGroupPrefix); err != nil {
		gtest.Error(err)
	} else {
		dbPrefix = r
	}

	// Invalid db.
	if r, err := gdb.NewByGroup(DBGroupInvalid); err != nil {
		gtest.Error(err)
	} else {
		dbInvalid = r
	}

	fmt.Println("init sqlite db finish")
}

func CreateTableFromSql() {
	// 初始化数据脚本
	sqlFile, err := os.Open(InitSql)
	if err != nil {
		panic(err)
	}
	defer sqlFile.Close()

	sqlBytes, err := ioutil.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	// 执行 SQL 脚本
	_, err = db.Exec(ctx, string(sqlBytes))
	if err != nil {
		panic(err)
	}
}

func DropAllTableWithDb() {
	fmt.Println("===========删除数据库")
	tables, err := db.Tables(ctx)
	if err != nil {
		gtest.Error(err)
	}
	for _, table := range tables {
		// SQLite 中的特殊表，用于存储自增列的当前值，无法删除
		if table == "sqlite_sequence" {
			continue
		}
		if _, err := db.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)); err != nil {
			gtest.Error(err)
		}
	}
}
