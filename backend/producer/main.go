package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GermanPachec0/producer/trades"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Failed to load enviroment")
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
