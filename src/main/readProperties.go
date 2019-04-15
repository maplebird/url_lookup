package main

import (
	"github.com/magiconair/properties"
	"log"
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
	if err := props.Decode(&dbConfig); err != nil {
		log.Fatal("FATAL: ", err)
	}

	return dbConfig
}

func getServerConfig(filename string) serverConfig {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var serverConfig serverConfig
	if err := props.Decode(&serverConfig); err != nil {
		log.Fatal("FATAL: ", err)
	}

	return serverConfig
}