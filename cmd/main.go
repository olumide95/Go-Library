package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/olumide95/go-library/bootstrap"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db := &bootstrap.Database{}
	db.InitDb()

}
