package main

import (
	"log"

	"github.com/guiaramos/bookstore/users/app"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.StartApplication()

}
