package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var TICKERS []string
var KAFKA_HOST string
var KAFKA_PORT string

func Load(envFilePath string) {
	if envFilePath != "" {
		// Load environment variables from the specified file
		err := godotenv.Load(envFilePath)
		if err != nil {
			log.Fatal("Failed to load environment file: ", err)
		}
	} else {
		// Load environment variables from the default file (e.g., .env in the same directory)
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to load environment file: ", err)
		}
	}

	// Rest of your existing code
	t := os.Getenv("TICKERS")
	TICKERS := strings.Split(t, ",")
	LoadTikers(TICKERS)

	KAFKA_HOST = os.Getenv("KAFKA_HOST")
	KAFKA_PORT = os.Getenv("KAFKA_PORT")
	fmt.Println(TICKERS)
}
