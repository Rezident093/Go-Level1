package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
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
	cfg := DBConfig{}

	flag.IntVar(&cfg.DBPort, "dbport", 8080, "port to connect")
	flag.StringVar(&cfg.DBUrl, "dburl", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Database connection URL")
	_, err = url.Parse("postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	if err != nil {
		log.Fatalf("failed parsing postgres DSN: %v", err)
	}
	flag.StringVar(&cfg.DBjaeger_url, "dbjaeger", "http://jaeger:16686", "Database connection URL")
	_, err = url.Parse("http://jaeger:16686")
	if err != nil {
		log.Fatalf("failed parsing postgres DSN: %v", err)
	}
	flag.StringVar(&cfg.DBsentry_url, "dbsentry", "http://http://sentry:9000", "Database connection URL")
	_, err = url.Parse("http://http://sentry:9000")
	if err != nil {
		log.Fatalf("failed parsing postgres DSN: %v", err)
	}
	flag.StringVar(&cfg.DBkafka_broker, "dbkafka", "kafka:9092", "Database connection broker address")
	flag.StringVar(&cfg.DBsome_app_id, "some_app_id", "testid", "Database connection")
	flag.StringVar(&cfg.DBsome_app_key, "some_app_key", "testkey", "Database connection keys")
	flag.Parse()

	fmt.Println(cfg)
}
