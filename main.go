package main

import (
	"chess/lichess"
	"context"
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

	account, err := lichessClient.GetAccount(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Authenticated as:", account.Username)

	game, err := lichessClient.CreateComputerGame(
		context.Background(),
		lichess.ComputerGameOptions{
			Level:          1,
			ClockLimit:     300,
			ClockIncrement: 0,
			Color:          "random",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Game created:", lichessClient.BaseURL+game.ID)
}
