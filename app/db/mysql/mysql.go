package mysql

import (
	"github.com/jinzhu/gorm"
	"myblog-api/app/config"
	"myblog-api/app/loger"

	//这一行需要保留，否则会报import _ "github.com/go-sql-driver/mysql"错误
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)


var MysqlDB *Mysql

type Mysql struct {
	conn *gorm.DB
}

func Default() {
	mysql := &Mysql{}
	conn, err := gorm.Open(config.Configs.DBDriver, config.Configs.DBUser+":"+config.Configs.DBPass+"@tcp("+config.Configs.DBHost+":"+config.Configs.DBPort+")/"+config.Configs.DBName)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		panic(err.Error() + config.Configs.DBDriver)
	}
	if config.Configs.DBDebug {
		conn = conn.Debug()
	}
	conn.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	conn.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	mysql.conn = conn
	MysqlDB = mysql
}

func (mysql *Mysql) GetConn() *gorm.DB {
	return mysql.conn
}
