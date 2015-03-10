package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Web struct {
		Port    int    `json:"port"`
		Address string `json:"address"`
	} `json:"web"`
	Cache struct {
		TTL     int    `json:"ttl"`
		Address string `json:"address"`
	} `json:"cache"`
	TMDb struct {
		APIKEY string `json:"api_key"`
	} `json:"tmdb"`
}

func Init() Configuration {
	var conf Configuration

	content, err := ioutil.ReadFile("wintergarten.conf")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &conf)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}
