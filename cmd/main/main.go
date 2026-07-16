package main

import (
	"entrywatchserver/internal/db"
	"entrywatchserver/internal/handlers"
	"entrywatchserver/internal/repository"
	"entrywatchserver/internal/router"
	"log"
	"net/http"
)


func main() {

	mongoClient,err := db.ConnectToMongo("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database("test")

	userRepo := repository.NewUserRepository(database)
	userHandler := handlers.NewUserHandler(userRepo)

	r := router.New(userHandler)

	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
