package main

import (
	"fmt"
	"log"

	"github.com/ArtemHvozdov/food-service-establishments/internal/db"
)

func main() {
	places, err := db.Places()
	if err != nil {
		log.Fatalf("generate: %v", err)
	}
	if err := db.Validate(places); err != nil {
		log.Fatalf("generate: валідація: %v", err)
	}

	fmt.Println(places)

	// TODO: script to generate pages (1.6-1.8)
}
