package main

import (
	"github.com/magiconair/properties"
	"log"
)

type dbConfig struct {
	host string `props:"dbprops.host,default=127.0.0.1"`
	port int `props:"dbprops.port,default=3306"`
	schema string `props:"dbprops.schema,default=url_lookup"`
	user string `props:"dbprops.user"`
	password string `props:"dbprops.password"`
}

func getDbConfig(filename string) dbConfig {
	props := properties.MustLoadFile(filename, properties.UTF8)

	var dbConfig dbConfig
	if err := props.Decode(&dbConfig); err != nil {
		log.Fatal(err)
	}

	return dbConfig
}

