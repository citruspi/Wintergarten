package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Web struct {
		Address string `json:"address"`
	} `json:"web"`
	Cache struct {
		TTL     int    `json:"ttl"`
		Address string `json:"address"`
		Enabled bool   `json:"enabled"`
	} `json:"cache"`
	TMDb struct {
		APIKEY string `json:"api_key"`
	} `json:"tmdb"`
}

func Init() Configuration {
	var conf Configuration
	var confPath string
	confPaths := []string{"/etc/wintergarten.conf", "/etc/wintergarten.json", "wintergarten.conf", "wintergarten.json"}

	for _, path := range confPaths {
		if _, err := os.Stat(path); err == nil {
			confPath = path
			break
		}
	}

	if confPath == "" {
		log.Fatal("No configuration found.")
	}

	content, err := ioutil.ReadFile(confPath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &conf)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}
