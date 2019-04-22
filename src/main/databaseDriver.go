package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Use connection string as global variable so properties file isn't read every time db connection is called
var dbConnectionString = getConnectionString()

// Parse properties file and create database connection string
func getConnectionString() (dbConnectionString string) {
	config := getConfig("config.properties")

	dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBSchema)

	return dbConnectionString
}

// Create new database connection that can be passed to different methods
func getDbConn() (db *sql.DB) {
	var err error
	db, err = sql.Open("mysql", dbConnectionString)
	fatalErr(err)
	return db
}

// Test database connection is working before starting the application
// Test database
// Expects at least 1 row in url_lookup.fqdns table
func testDbConn(db *sql.DB) bool {
	var fqdn string
	var reputation string

	log.Println("Initializing new database connection")
	err := db.QueryRow("SELECT * FROM fqdns LIMIT 1").Scan(&fqdn, &reputation)
	fatalErr(err)

	err = db.Close()
	fatalErr(err)

	log.Println("Selected 1 row from database: ", fqdn, " ", reputation)
	if fqdn != "" {
		return true
	}
	return false
}
