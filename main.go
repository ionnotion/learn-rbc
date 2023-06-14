package main

import (
	"fmt"
	"learn-rbc/db"
	"learn-rbc/router"
	"log"
)

func main() {
	fmt.Println("Hello World!")

	database, err := db.Open()
	if err != nil {
		log.Fatalf("Database Connection failure: %s", err)
	}

	log.Print(database)
	router.Init()
}
