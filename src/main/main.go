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

	lookupUrl(requestedUrl)
}

func lookupUrl(requestedUrl string) string {
	fmt.Println("Looking up this URL: ", requestedUrl)
	return requestedUrl
}

func main() {
	dbConfig := getDbConfig("config.properties")

	fmt.Println("DB host: ", dbConfig.host)
	fmt.Println("DB port: ", dbConfig.port)
	fmt.Println("DB schema: ", dbConfig.schema)
	fmt.Println("DB user: ", dbConfig.user)

	http.HandleFunc("/1/urlinfo/", parseRequest)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}