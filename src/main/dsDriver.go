package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/magiconair/properties"
	"log"
)

type dbConfig struct {
	Host     string `props:"host,default=127.0.0.1"`
	Port     string `props:"port,default=3306"`
	Schema   string `props:"schema,default=url_lookup"`
	User     string `props:"user"`
	Password string `props:"password"`
}

func getDbConfig(filename string) dbConfig {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var dbConfig dbConfig
	if err := props.Decode(&dbConfig); err != nil {
		log.Fatal("FATAL: ", err)
	}

	return dbConfig
}

func getConnString() (connString string) {
	dbConfig := getDbConfig("db.properties")

	log.Println("Initializing new database connection")
	log.Printf("DB host: %s:%s\n", dbConfig.Host, dbConfig.Port)
	log.Println("DB schema: ", dbConfig.Schema)

	connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	return connString
}

func dbInit() {
	var db *sql.DB
	var err error
	connString := getConnString()
	db, err = sql.Open("mysql", connString)
	checkErr(err)

	// Test database
	// Expects at least 1 row in url_lookup.fqdns table
	rows, err := db.Query("SELECT * FROM fqdns LIMIT 1")
	checkErr(err)

	for rows.Next() {
		var fqdn string
		var reputation string
		err = rows.Scan(&fqdn, &reputation)
		fmt.Println(fqdn, reputation)
	}

	db.Close()
}

func dbQuery(query string) (queryResult string) {
	var db *sql.DB
	var err error
	connString := getConnString()
	db, err = sql.Open("mysql", connString)
	checkErr(err)

	rows, err := db.Query(query)
	checkErr(err)

	db.Close()

	for rows.Next() {
		var fqdn string
		var reputation string
		err = rows.Scan(&fqdn, &reputation)
		fmt.Println(fqdn, reputation)
	}
}








