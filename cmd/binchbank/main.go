package main

import (
	"flag"
	"log"

	"github.com/estromenko/binchbank/app"
	"github.com/joho/godotenv"
)

var (
	dotenvPath = flag.String("dotenv", ".env", "Path to .env file")
)

func main() {
	flag.Parse()

	if err := godotenv.Load(*dotenvPath); err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}
