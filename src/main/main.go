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

// Fatal error
func fatalErr(err error) {
	if err != nil {
		log.Fatal("FATAL ERROR: ", err, "\nEXITING url_lookup")
	}
}

// Log errors for non-critical components or general runtime issues
func logErr(err error) {
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
