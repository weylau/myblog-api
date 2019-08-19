package db

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
)

func DBConn() (db *gorm.DB) {
	dbDriver := "mysql"
	dbHost := "172.16.57.110"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "123456"
	dbName := "myblog"
	dbDebug := true

	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	if dbDebug {
		db = db.Debug()
	}
	return db
}
