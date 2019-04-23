package main

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

type config struct {
	DBHost        string
	DBPort        string
	DBSchema      string
	DBUser        string
	DBPassword    string
	ServerAddress string
	ServerPort    string
}

func getConfig(filename string) config {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var config config
	err := props.Decode(&config)
	fatalErr(err)

	// Override any values set in config.properties with environment variables
	config = setConfigFromEnvVars(config)

	return config
}

func setConfigFromEnvVars(config config) config {
	// There is probably a better way to do this, like looping over a map

	if os.Getenv("URL_LOOKUP_DBHOST") != "" {
		config.DBHost = os.Getenv("URL_LOOKUP_DBHOST")
		log.Println("Overriding database host from env var with value ", os.Getenv("URL_LOOKUP_DBHOST"))
	}
	if os.Getenv("URL_LOOKUP_DBPORT") != "" {
		config.DBPort = os.Getenv("URL_LOOKUP_DBPORT")
		log.Println("Overriding database port from env var with value ", os.Getenv("URL_LOOKUP_DBPORT"))
	}
	if os.Getenv("URL_LOOKUP_DBSCHEMA") != "" {
		config.DBSchema = os.Getenv("URL_LOOKUP_DBSCHEMA")
		log.Println("Overriding database schema from env var with value ", os.Getenv("URL_LOOKUP_DBSCHEMA"))
	}
	if os.Getenv("URL_LOOKUP_DBUSER") != "" {
		config.DBUser = os.Getenv("URL_LOOKUP_DBUSER")
		log.Println("Overriding database user from env var with value ", os.Getenv("URL_LOOKUP_DBUSER"))
	}
	if os.Getenv("URL_LOOKUP_DBPASSWORD") != "" {
		config.DBPassword = os.Getenv("URL_LOOKUP_DBPASSWORD")
		log.Println("Overriding database password from env var with value ", os.Getenv("URL_LOOKUP_DBPASSWORD"))
	}
	if os.Getenv("URL_LOOKUP_SERVER_ADDRESS") != "" {
		config.ServerAddress = os.Getenv("URL_LOOKUP_SERVER_ADDRESS")
		log.Println("Overriding server listen address from env var with value ", os.Getenv("URL_LOOKUP_SERVER_ADDRESS"))
	}
	if os.Getenv("URL_LOOKUP_SERVER_PORT") != "" {
		config.ServerPort = os.Getenv("URL_LOOKUP_SERVER_PORT")
		log.Println("Overriding server listen port from env var with value ", os.Getenv("URL_LOOKUP_SERVER_PORT"))
	}

	return config
}
