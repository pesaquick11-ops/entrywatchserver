package main

import (
	"entrywatchserver/internal/db"
	"entrywatchserver/internal/handlers"
	"entrywatchserver/internal/repository"
	"entrywatchserver/internal/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("Mongo DB URI is not set")
	}

	mongoClient, err := db.ConnectToMongo(mongoURI)

	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database("test")

	userRepo := repository.NewUserRepository(database)
	userHandler := handlers.NewUserHandler(userRepo)

	r := router.New(userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	addr := ":" + port

	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
