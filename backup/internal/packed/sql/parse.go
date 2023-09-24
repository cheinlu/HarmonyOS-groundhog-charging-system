package sql

import (
	"fmt"
	"strings"
)

// SQLParser 结构体用于解析和修改 SQL 语句
type SQLParser struct {
	sql  string
	args []interface{}
}

// NewSQLParser 创建一个新的 SQLParser 实例
func NewSQLParser(sql string, args []interface{}) *SQLParser {
	return &SQLParser{
		sql:  sql,
		args: args,
	}
}

var SqlTokens = []string{"group", "order", "limit"}

// AddWhere 方法用于向 SQL 语句中添加 WHERE 条件
// 它接受一个条件和一个值，然后将它们插入到 SQL 语句的适当位置
//
// 示例：
// parser := NewSQLParser("select * from user join order on user.id = order.user_id", nil)
// parser.AddWhere("tenant_id", 1)
func (p *SQLParser) AddWhere(condition string, value interface{}) *SQLParser {
	whereSQL := fmt.Sprintf("%s=?", condition)
	lowerSQL := strings.ToLower(p.sql)
	index := -1
	if strings.Contains(lowerSQL, "where") {
		whereSQL = fmt.Sprintf(" %s and", whereSQL)
		index = strings.Index(lowerSQL, " where ") + len(" where")

	} else if strings.Contains(lowerSQL, " group ") || strings.Contains(lowerSQL, " order ") || strings.Contains(lowerSQL, " limit ") {
		whereSQL = fmt.Sprintf(" where %s", whereSQL)
		groupIndex := strings.Index(lowerSQL, " group ")
		orderIndex := strings.Index(lowerSQL, " order ")
		limitIndex := strings.Index(lowerSQL, " limit ")
		index = groupIndex
		if orderIndex != -1 && (index == -1 || index > orderIndex) {
			index = orderIndex
		}
		if limitIndex != -1 && (index == -1 || index > limitIndex) {
			index = limitIndex
		}
	} else {
		whereSQL = fmt.Sprintf(" where %s", whereSQL)
		index = len(p.sql)
	}
	p.sql = p.insertCondition(lowerSQL, whereSQL, index)

	newArgs := make([]interface{}, len(p.args)+1)
	newArgs[0] = value
	copy(newArgs[1:], p.args)
	p.args = newArgs
	return p
}

// insertCondition 方法用于在 SQL 语句的下标后面插入条件
// 它接受一个 SQL 语句、一个条件和一个关键字，然后返回修改后的 SQL 语句
//
// 示例：
// p.sql = p.insertCondition(lowerSQL, whereSQL, 10)
func (p *SQLParser) insertCondition(lowerSQL, condition string, index int) string {
	// index := strings.Index(lowerSQL, keyword)
	preSQL := p.sql[:index]
	afterSQL := p.sql[index:]
	return preSQL + condition + afterSQL
}

// GetSQL 方法返回当前的 SQL 语句
//
// 示例：
// fmt.Println(parser.GetSQL())
func (p *SQLParser) GetSQL() string {
	return p.sql
}

// GetArgs 方法返回当前的参数列表
//
// 示例：
// fmt.Println(parser.GetArgs())
func (p *SQLParser) GetArgs() []interface{} {
	return p.args
}
