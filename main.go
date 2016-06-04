package main

import (
	"log"
	"os"
)

func main() {
	token := os.Getenv("GITTER_TOKEN")

	if token == "" {
		log.Fatal("GITTER_TOKEN environment variable is required")
	}

	bot := NewBot(token)
	bot.AddHandler(&HeardYou{})
	bot.AddHandler(&Ping{})
	bot.Listen()
}
