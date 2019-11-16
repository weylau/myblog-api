package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/weylau/myblog-api/app/config"
)

type Mysql struct {
	conn *gorm.DB
}

func Default() (db *Mysql) {
	mysql := &Mysql{}
	conn, err := gorm.Open(config.Configs.DBDriver, config.Configs.DBUser+":"+config.Configs.DBPass+"@tcp("+config.Configs.DBHost+":"+config.Configs.DBPort+")/"+config.Configs.DBName)
	if err != nil {
		panic(err.Error())
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
