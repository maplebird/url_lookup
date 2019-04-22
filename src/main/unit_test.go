package main

import (
	"os"
	"strings"
	"testing"
)

// Unit tests below
func Test_getConfig(t *testing.T) {
	t.Log("TEST 1 - parsing config.properties")

	config := getConfig("config.properties")
	if config.DBHost != ""{
		t.Log("PASS - successfully retrieved DBHost configuration property")
	} else {
		t.Error("FAIL - unable to retrieve DBHost configuration property")
	}
}

func Test_setConfigFromEnvVars(t *testing.T) {
	t.Log("TEST 2 - reading config from env vars")

	var config config
	os.Setenv("URL_LOOKUP_DBHOST", "localhost")
	config = setConfigFromEnvVars(config)

	if config.DBHost == "localhost" {
		t.Log("PASS - successfully set DBHost property from env vars")
	} else {
		t.Error(("FAIL - unable to set DBHost property from env vars"))
	}
}

func Test_getConnectionString(t *testing.T) {
	t.Log("TEST 3 - get database connection string")

	dbConnectionString = getConnectionString()

	t.Log("Database connection string currently set to:")
	t.Log(dbConnectionString)

	indexOfDbUser := strings.IndexAny(dbConnectionString, "localhost")

	if indexOfDbUser >= 0 {
		t.Log("PASS - database connection string contains correct database host")
	} else {
		t.Error("FAIL - db connection string does not have correct database host")
	}
}

func Test_parseRequestedUrl_fqdnOnly(t *testing.T) {
	t.Log("TEST 4 - parse requested URL (only FQDN given, no path)")

	parsedUrl := parseRequestedUrl("www.google.com")

	if parsedUrl.fqdn == "www.google.com" && parsedUrl.path == "" {
		t.Log("PASS - parsing returns fqdn and no path")
	} else {
		t.Error("FAIL - parsing does not return FQDN or returns non-empty path")
	}
}

func Test_parseRequestedUrl_fullPath(t *testing.T) {
	t.Log("TEST 5 - parse requested URL with path")

	parsedUrl := parseRequestedUrl("www.google.com/foo/bar")

	if parsedUrl.fqdn == "www.google.com" && parsedUrl.path == "/foo/bar" {
		t.Log("PASS - parsing returns correct fqdn and path")
	} else {
		t.Error("FAIL - parsing does not return correct FQDN or path")
	}
}



