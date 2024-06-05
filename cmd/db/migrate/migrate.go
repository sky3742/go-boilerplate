package main

import (
	"go-boilerplate/cmd/db"
	"log"
)

func main() {
	err := db.Migrations.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
