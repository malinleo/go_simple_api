package main

import (
	"log"
	"simple_api/src/db"
)

func main() {
	if dbInitErr := db.Initialize(); dbInitErr != nil {
		log.Fatalf("bacbacbacb: %s", dbInitErr.Error())
	}
}