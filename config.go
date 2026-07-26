package main

import (
	"log"
	"os"
)

type Config struct {
	token string
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Missing required env var: %s", key)
	}
	return v
}
