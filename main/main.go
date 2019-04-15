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
	var requested_url string = strings.Replace(path, "/1/urlinfo/", "", 1)

	lookupUrl(requested_url)
}

func lookupUrl(requested_url string) string {
	fmt.Println("Looking up this URL: ", requested_url)
	return requested_url
}

func main() {
	http.HandleFunc("/1/urlinfo/", parseRequest)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}