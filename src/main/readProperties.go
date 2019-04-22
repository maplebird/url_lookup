package main

import (
	"github.com/magiconair/properties"
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

func getEnvVars(envVar string) string {
	return os.Getenv(envVar)
}

func getConfig(filename string) config {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var config config
	err := props.Decode(&config)
	fatalErr(err)

	return config
}
