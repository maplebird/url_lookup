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

	connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	return connString
}

// Create new database connection that can be passed to different methods
func getDbConn() (db *sql.DB) {
	var err error
	connString := getConnString()
	db, err = sql.Open("mysql", connString)
	fatalErr(err)
	return db
}

// Test database connection is working before starting the application
// Test database
// Expects at least 1 row in url_lookup.fqdns table
func testDbConn(db *sql.DB) (bool) {
	var fqdn string
	var reputation string

	log.Println("Initializing new database connection")
	err := db.QueryRow("SELECT * FROM fqdns LIMIT 1").Scan(&fqdn, &reputation)
	fatalErr(err)

	db.Close()

	log.Println("Selected 1 row from database: ", fqdn, reputation)
	if fqdn != "" {
		return true
	}
	return false
}





