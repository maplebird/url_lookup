package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getConnString() (connString string) {
	dbConfig := getDbConfig("config.properties")

	log.Println("Initializing new database connection")
	log.Printf("DB host: %s:%s\n", dbConfig.Host, dbConfig.Port)
	log.Println("DB schema: ", dbConfig.Schema)

	connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	return connString
}

func getDbConn() (db *sql.DB) {
	var err error
	connString := getConnString()
	db, err = sql.Open("mysql", connString)
	checkErr(err)

	return db


}

//func dbQuery(query string) (queryResult string) {
//	var db *sql.DB
//	var err error
//	connString := getConnString()
//	db, err = sql.Open("mysql", connString)
//	checkErr(err)
//
//	rows, err := db.Query(query)
//	checkErr(err)
//
//	db.Close()
//
//	for rows.Next() {
//		var fqdn string
//		var reputation string
//		err = rows.Scan(&fqdn, &reputation)
//		fmt.Println(fqdn, reputation)
//	}
//}


func testDbConn(db *sql.DB) (bool) {
	var output string

	//Test database
	//Expects at least 1 row in url_lookup.fqdns table
	rows, err := db.Query("SELECT * FROM fqdns LIMIT 1")
	checkErr(err)

	for rows.Next() {
		var fqdn string
		var reputation string
		err = rows.Scan(&fqdn, &reputation)
		output = fqdn + reputation
	}

	if output != "" {
		return true
	}
	return false
}






