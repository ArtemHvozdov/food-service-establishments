package main

import (
	"fmt"

	"github.com/ArtemHvozdov/food-service-establishments/internal/db"
)

func main() {
	places := db.Places()

	fmt.Println(places)

	// TODO: script to validate places and generate pages
}
