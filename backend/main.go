package main

import (
	"log"
	"mockwallet/auth"
	"mockwallet/db"
	"os"
)

func main() {
	auth.SecretKey = []byte(os.Getenv("JWT_SECRET"))

	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	} //connect database to handlers

}
