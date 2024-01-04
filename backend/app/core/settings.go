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

func Load() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load environment file")
	}
	t := os.Getenv("TICKERS")
	TICKERS := strings.Split(t, ",")
	LoadTikers(TICKERS)

	KAFKA_HOST = "127.0.0.1"
	KAFKA_PORT = "9092"
	fmt.Println(TICKERS)
}
