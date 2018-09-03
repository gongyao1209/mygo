package db

import (
	"database/sql"
	"mygo/config"
	"fmt"

	_"github.com/go-sql-driver/mysql"
)

var db_ *sql.DB

func Get() *sql.DB {
	return db_
}

type DbWorker struct {
	Dsn	string
	Db	*sql.DB
}

func init()  {
	var err error

	dbw := DbWorker{
		Dsn:config.Get("Mysql"),
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		fmt.Printf("SQL ERROR")
		return
	}

	db_ = dbw.Db
}