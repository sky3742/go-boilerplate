package main

import (
	"go-boilerplate/cmd/db"
	"log"
)

func main() {
	err := db.Seeds.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
