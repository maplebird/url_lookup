package main

import (
	"github.com/magiconair/properties"
)

type dbConfig struct {
	Host     string
	Port     string
	Schema   string
	User     string
	Password string
}

type serverConfig struct {
	ServerAddress string
	ServerPort string
}

func getDbConfig(filename string) dbConfig {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var dbConfig dbConfig
	err := props.Decode(&dbConfig)
	fatalErr(err)

	return dbConfig
}

func getServerConfig(filename string) serverConfig {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var serverConfig serverConfig
	err := props.Decode(&serverConfig)
	fatalErr(err)

	return serverConfig
}