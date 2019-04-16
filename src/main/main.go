package main

import (
	"log"
)

func main() {
	var db = getDbConn()

	// Test database connection before starting server
	if testDbConn(db){
		log.Println("Successfully connected to database")
	}else {
		log.Fatal("ERROR: Cannot connect to databse. Exiting.")
	}

	router()
}

// Fatal error, exit program
func fatalErr(err error) {
	if err != nil {
		log.Fatal("FATAL ERROR: ", err, "\nEXITING url_lookup")
	}
}

// Non-fatal error, log only
func logErr(err error) {
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
