package main

import (
	"log"
	"mockwallet/auth"
	"mockwallet/db"
	"mockwallet/handlers"
	"mockwallet/middleware"
	"net/http"
	"os"
)

func main() {
	auth.SecretKey = []byte(os.Getenv("JWT_SECRET"))

	_, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/balance", middleware.AuthMiddleware(handlers.GetBalance))

	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, middleware.WithCors(mux)))
}
