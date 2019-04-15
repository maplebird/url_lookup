package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Parse properties file and create database connection string
func getConnString() (connString string) {
	dbConfig := getDbConfig("config.properties")

	log.Println("Initializing new database connection")
	log.Printf("DB host: %s:%s\n", dbConfig.Host, dbConfig.Port)
	log.Println("DB schema: ", dbConfig.Schema)

	connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	return connString
}

// Create new database connection that can be passed to different methods
func getDbConn() (db *sql.DB) {
	var err error
	connString := getConnString()
	db, err = sql.Open("mysql", connString)
	checkErr(err)
	return db
}

// Test database connection is working before starting the application
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






