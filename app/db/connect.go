package db

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbHost := "172.16.57.110"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "123456"
	dbName := "myblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
