package main

import (
	"log"

	"github.com/cartman720/go-whisper/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
