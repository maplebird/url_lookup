package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type response struct {
	RequestedUrl string
	Reputation string
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

	fmt.Fprint(w, createResponse(requestedUrl,  reputation))
}

func createResponse(requestedUrl string, reputation string) (restResponse string) {
	output := response{
		RequestedUrl: requestedUrl,
		Reputation: reputation,
	}

	restResponseEncoded, err := json.Marshal(output)
	logErr(err)

	return string(restResponseEncoded)
}


// Start web server
func router() {
	serverConfig := getServerConfig("config.properties")
	listenAddress := fmt.Sprintf("%s:%s", serverConfig.ServerAddress, serverConfig.ServerPort)

	log.Println("Initializing URL Lookup server")
	log.Println(fmt.Sprintf("Listening on %s address (will listen on all interfaces if empty)", serverConfig.ServerAddress))
	log.Println("Listening on port ", serverConfig.ServerPort)

	http.HandleFunc("/1/urlinfo/", processRequest)
	err := http.ListenAndServe(listenAddress, nil)
	fatalErr(err)
}