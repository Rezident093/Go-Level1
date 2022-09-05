package main

import (
	"fmt"
	"module/config"
)

type DBConfig struct {
	DBPort         int
	DBUrl          string
	DBjaeger_url   string
	DBsentry_url   string
	DBkafka_broker string
	DBsome_app_id  string
	DBsome_app_key string
}

func main() {
	var err error
	cfg, err := config.ConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
