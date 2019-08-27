package db

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
	"github.com/weylau/myblog-api/app/configs"
)

func DBConn() (db *gorm.DB) {

	db, err := gorm.Open(configs.Configs.DBDriver, configs.Configs.DBUser+":"+configs.Configs.DBPass+"@tcp("+configs.Configs.DBHost+":"+configs.Configs.DBPort+")/"+configs.Configs.DBName)
	if err != nil {
		panic(err.Error())
	}
	if configs.Configs.DBDebug {
		db = db.Debug()
	}
	return db
}
