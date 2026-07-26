package main

import (
	"chess/lichess"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env:", err)
	}
	requireEnv("LICHESS_PAT")

	cfg := &Config{
		token: os.Getenv("LICHESS_PAT"),
	}

	lichessClient := lichess.NewLichessClient(
		cfg.token,
	)
	fmt.Print(lichessClient)
	fmt.Println("Created client successfully")
}
