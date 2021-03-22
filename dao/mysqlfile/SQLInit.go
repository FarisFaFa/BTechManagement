package mysqlfile

import (
	"database/sql"
	"fmt"

	"github.com/go-xorm/xorm"
)

var db *sql.DB
var x *xorm.Engine

// 数据库初始化
func InitDB() (err error) {
	// 配置DSN，Data Source Name
	dsn := "root:WJT-19930208wjt@tcp(18.191.157.203:3306)/BTech"
	// 用什么数据库
	db, err = sql.Open("mysql", dsn)
	// 查看是否能从dsn中提取到数据库信息，如不能，打印错误
	if err != nil {
		fmt.Printf("error is %v", err)
	}
	// 尝试连接数据库，如无法连接，打印错误
	err = db.Ping()
	if err != nil {
		fmt.Printf("error is %v", err)
	}
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	// 设置最大连接数
	db.SetMaxOpenConns(10)
	return
}

// xorm初始化
func XormInit() (err error) {
	dsn := "root:WJT-19930208wjt@tcp(18.191.157.203:3306)/BTech"
	x, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败，错误为%v", err)
	}
	x.ShowSQL(true)
	return
}
