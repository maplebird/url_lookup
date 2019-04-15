package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func dbInit(connString string) {
	sql.Open("mysql", connString)
}