package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type candidateUrl struct {
	fqdn string
	path string
}

func parseRequestedUrl(requestedUrl string) candidateUrl {
	var parsedUrl candidateUrl

	var urlPathSeparatorIndex int
	urlPathSeparatorIndex = strings.IndexAny(requestedUrl, "/")

	if urlPathSeparatorIndex > 0 {
		parsedUrl.fqdn = requestedUrl[:urlPathSeparatorIndex]
		parsedUrl.path = requestedUrl[urlPathSeparatorIndex:]
	} else {
		parsedUrl.fqdn = requestedUrl
		parsedUrl.path = ""
	}

	return parsedUrl
}

func lookupUrl(requestedUrl string) (reputation string) {
	var parsedUrl candidateUrl
	parsedUrl = parseRequestedUrl(requestedUrl)
	reputation = checkReputation(parsedUrl)
	return reputation
}

func checkReputation(parsedUrl candidateUrl) (reputation string) {
	var db = getDbConn()
	var query = fmt.Sprintf("SELECT reputation FROM fqdns WHERE fqdn = '%s'", parsedUrl.fqdn)

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
		reputation = checkMixedReputation(parsedUrl, db)
	}

	err = db.Close()
	logErr(err)

	return reputation
}

func checkMixedReputation(parsedUrl candidateUrl, db *sql.DB) (reputation string) {
	var query = fmt.Sprintf(
		"SELECT reputation FROM path_lookup WHERE fqdn = '%s' AND path = '%s'",
		parsedUrl.fqdn, parsedUrl.path)

	err := db.QueryRow(query).Scan(&reputation)
	if err == sql.ErrNoRows {
		log.Println("Full URL path not found in database, domain reputation mixed.")
		reputation = "mixed"
	} else {
		log.Println("Full URL reputation: ", reputation)
	}

	err = db.Close()
	logErr(err)

	return reputation
}