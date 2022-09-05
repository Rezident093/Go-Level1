package config

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type DBConfig struct {
	// port to connect
	DBPort int `json:"db_port"`
	// Database connection URL
	DBUrl string `json:"db_url"`
	// Database connection URL
	DBjaeger_url string `json:"db_jaeger"`
	// Database connection URL
	DBsentry_url string `json:"db_sentry"`
	// Database connection broker address
	DBkafka_broker string `json:"db_kafka"`
	// Database connection
	DBsome_app_id string `json:"some_app_id"`
	// Database connection keys
	DBsome_app_key string `json:"some_app_key"`
}

func ConfigFromFile(filename string) (*DBConfig, error) {
	plan, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := &DBConfig{}
	err = json.Unmarshal(plan, conf)
	if err != nil {
		return nil, err
	}

	if _, err := url.Parse(conf.DBUrl); err != nil {
		return nil, err
	}

	if _, err := url.Parse(conf.DBjaeger_url); err != nil {
		return nil, err
	}

	if _, err := url.Parse(conf.DBsentry_url); err != nil {
		return nil, err
	}

	return conf, nil
}
