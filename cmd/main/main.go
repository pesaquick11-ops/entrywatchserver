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
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func setupHandlers(database *mongo.Database) router.Handlers {
	userRepo := repository.NewUserRepository(database)

	return router.Handlers{
		User: handlers.NewUserHandler(userRepo),
	}
}

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

	dbName := os.Getenv("DATABASE_NAME")

	database := mongoClient.Database(dbName)

	h := setupHandlers(database)
	r := router.New(h)

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
