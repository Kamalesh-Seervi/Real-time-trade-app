package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	trades "producer/trades"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Failed to Load ENV file...")
	}

	t := os.Getenv("TICKERS")
	topics := strings.Split(t, ",")
	for i, topic := range topics {
		topics[i] = strings.Trim(strings.Trim(topic, "\\"), "\"")
	}

	trades.SubScribeAndListen(
		topics,
	)
}
