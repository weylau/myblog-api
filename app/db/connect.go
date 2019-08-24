package db

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
	"github.com/weylau/myblog-api/app/configs"
)

func DBConn() (db *gorm.DB) {

	db, err := gorm.Open(configs.DBDriver, configs.DBUser+":"+configs.DBPass+"@tcp("+configs.DBHost+":"+configs.DBPort+")/"+configs.DBName)
	if err != nil {
		panic(err.Error())
	}
	if configs.DBDebug {
		db = db.Debug()
	}
	return db
}
