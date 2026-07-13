package main

import (
	"log"

	"github.com/ArtemHvozdov/food-service-establishments/internal/db"
	"github.com/ArtemHvozdov/food-service-establishments/internal/render"
)

func main() {
	places, err := db.Places()
	if err != nil {
		log.Fatalf("generate: %v", err)
	}
	if err := db.Validate(places); err != nil {
		log.Fatalf("generate: валідація: %v", err)
	}

	if err := render.Generate(places, "public"); err != nil {
		log.Fatalf("generate: рендер: %v", err)
	}
}
