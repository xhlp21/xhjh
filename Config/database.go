package Config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// SqlDB 全局mysql 连接器 DB
var SqlDB *sql.DB

// 定义连接数据库配置
type dbConfig struct {
	dbType      string //数据库类型 mysql
	dbName      string //数据库名称
	user        string //数据库账号/用户名
	password    string //数据库密码
	host        string //数据库ip 域名:端口
	tablePrefix string //表头
}

// 自动调用方法
func init() {
	//定义错误信息
	var err error
	//配置信息
	cfg := &dbConfig{
		dbType:      "mysql",
		dbName:      "xhj_gin",
		user:        "root",
		password:    "root",
		host:        "127.0.0.1:3306",
		tablePrefix: "",
	}
	//连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.user,
		cfg.password,
		cfg.host,
		cfg.dbName)

	//打开链接
	SqlDB, err = sql.Open("mysql", dsn)

	//错误信息提示
	if err != nil {
		log.Fatal("数据库连接错误:", err)
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	//设置空闲连接池中的最大连接数
	SqlDB.SetMaxIdleConns(20)
	//设置空闲连接池中的最大连接数
	SqlDB.SetMaxOpenConns(20)
}
