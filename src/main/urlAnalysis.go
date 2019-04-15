package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func lookupUrl(requestedUrl string) (reputation string) {
	var fqdn string
	var path string

	var if_has_path int
	if_has_path = strings.IndexAny(requestedUrl, "/")
	fmt.Println(if_has_path)

	if if_has_path > 0 {
		fqdn = requestedUrl[:if_has_path]
		path = requestedUrl[if_has_path:]
	} else {
		fqdn = requestedUrl
		path = ""
	}

	return checkReputation(fqdn, path)
}

func checkReputation(fqdn string, path string) (reputation string) {
	var db = getDbConn()
	var query = fmt.Sprintf("SELECT reputation FROM fqdns WHERE fqdn = '%s'", fqdn)

	err := db.QueryRow(query).Scan(&reputation)
	if err == sql.ErrNoRows {
		reputation = "unknown"
		log.Println("URL not found in DB, reputation is unknown")
	}

	switch reputation {
	case "safe":
		log.Println("Reputation is safe")
	case "unsafe":
		log.Println("URL is unsafe")
	case "mixed":
		// In case of mixed reputation, need to check both domain and path to the object being retrieved
		log.Println("Reputation for this domain is mixed, checking path_lookup table")
		reputation = checkMixedReputation(fqdn, path, db)
	}

	db.Close()

	return reputation
}

func checkMixedReputation(fqdn string, path string, db *sql.DB) (reputation string) {
	var query = fmt.Sprintf("SELECT reputation FROM path_lookup WHERE fqdn = '%s' AND path = '%s'", fqdn, path)
	log.Println(query)

	err := db.QueryRow(query).Scan(&reputation)
	if err == sql.ErrNoRows {
		log.Println("Full URL path not found in database, domain reputation mixed.")
		reputation = "mixed"
	} else {
		log.Println("Full URL reputation: ", reputation)
	}

	return reputation
}