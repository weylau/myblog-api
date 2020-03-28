package mysql

import (
	"github.com/jinzhu/gorm"
	"myblog-api/app/config"
	"myblog-api/app/loger"

	//这一行需要保留，否则会报import _ "github.com/go-sql-driver/mysql"错误
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)

type Mysql struct {
	conn *gorm.DB
}

func Default() (db *Mysql) {
	mysql := &Mysql{}
	conn, err := gorm.Open(config.Configs.DBDriver, config.Configs.DBUser+":"+config.Configs.DBPass+"@tcp("+config.Configs.DBHost+":"+config.Configs.DBPort+")/"+config.Configs.DBName)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		panic(err.Error() + config.Configs.DBDriver)
	}
	if config.Configs.DBDebug {
		conn = conn.Debug()
	}
	mysql.conn = conn
	return mysql
}

func (this *Mysql) GetConn() *gorm.DB {
	return this.conn
}
