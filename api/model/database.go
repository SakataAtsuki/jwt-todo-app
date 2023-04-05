package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func DbConnect() {
	dsn := "root:root@tcp(mysql:3307)/tododb?charset=utf8&parseTime=true"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}
