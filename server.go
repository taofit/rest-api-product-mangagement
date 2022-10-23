package main

import (
	"fmt"
	"log"

	"github.com/taofit/rest-api-product-mangagement/db"
)

func main() {
	db, err := db.SetDbCon()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	fmt.Println(db)
}
