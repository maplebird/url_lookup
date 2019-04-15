package main

import (
	"fmt"
	"net/http"
	"log"
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

	query := fmt.Sprint(`SELECT reputation FROM fqdns WHERE fqdn = "%s"`, requestedUrl)

	reputation = dbQuery(query)

	return requestedUrl
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	dbInit()

	http.HandleFunc("/1/urlinfo/", parseRequest)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}