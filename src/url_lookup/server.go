package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type urlInfo struct {
	RequestedUrl string
	Reputation   string
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	var requestedUrl string
	var reputation string

	err := r.ParseForm()
	logErr(err)

	path := r.URL.Path
	requestedUrl = strings.Replace(path, "/1/urlinfo/", "", 1)

	log.Println("Looking up URL: ", requestedUrl)
	reputation = lookupUrl(requestedUrl)

	fmt.Fprint(w, createResponse(requestedUrl, reputation))
}

func createResponse(requestedUrl string, reputation string) (restResponse string) {
	output := urlInfo{
		RequestedUrl: requestedUrl,
		Reputation:   reputation,
	}

	restResponseEncoded, err := json.Marshal(output)
	logErr(err)

	return string(restResponseEncoded)
}

// Start web server
func router() {
	config := getConfig("config.properties")
	listenAddress := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)

	log.Println("Initializing URL Lookup server")
	log.Println(fmt.Sprintf("Listening on %s address (will listen on all interfaces if empty)", config.ServerAddress))
	log.Println("Listening on port ", config.ServerPort)

	http.HandleFunc("/1/urlinfo/", processRequest)
	err := http.ListenAndServe(listenAddress, nil)
	fatalErr(err)
}
