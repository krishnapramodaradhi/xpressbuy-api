package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/krishnapramodaradhi/xpressbuy-api/internal/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("there was an error loading the env variables")
	}
}

func main() {
	s := config.NewServer(":8443")
	s.Run()
}
