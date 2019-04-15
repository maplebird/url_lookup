package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func parseRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := r.URL.Path
	var requestedUrl string = strings.Replace(path, "/1/urlinfo/", "", 1)

	log.Println("Looking up URL: ", requestedUrl)
	lookupUrl(requestedUrl)
}

func lookupUrl(requestedUrl string) (reputation string) {
	fmt.Println("Looking up this URL: ", requestedUrl)

	//query := fmt.Sprint(`SELECT reputation FROM fqdns WHERE fqdn = "%s"`, requestedUrl)

	//reputation = dbQuery(query)

	return requestedUrl
}

// Basic error handling function
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Start web server
func router() {
	serverConfig := getServerConfig("config.properties")
	listenAddress := fmt.Sprintf("%s:%s", serverConfig.ServerAddress, serverConfig.ServerPort)

	log.Println("Initializing URL Lookup server")
	log.Println("Listening on %s address (will listen on all interfaces if empty)", serverConfig.ServerAddress)
	log.Println("Listening on port ", serverConfig.ServerPort)

	http.HandleFunc("/1/urlinfo/", parseRequest)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


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