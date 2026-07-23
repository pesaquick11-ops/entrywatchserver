package main

import (
	"entrywatchserver/internal/db"
	"entrywatchserver/internal/handlers"
	"entrywatchserver/internal/mqtt"
	"entrywatchserver/internal/repository"
	"entrywatchserver/internal/router"
	"entrywatchserver/internal/ws"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func setupHandlers(database *mongo.Database, hub *ws.Hub) router.Handlers {
	userRepo := repository.NewUserRepository(database)
	attRepo := repository.NewAttendanceRepository(database)
	return router.Handlers{
		User:  handlers.NewUserHandler(userRepo),
		Athan: handlers.NewAttendanceHandler(attRepo),
		Hub:   hub,
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

	// --- MQTT + WebSocket telemetry setup ---
	hub := ws.NewHub()

	mqttCfg := mqtt.Config{
		Broker:   os.Getenv("MQTT_BROKER"),
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
		ClientID: "entrywatch-server",
		Topic:    "images/new", // or make this an env var too
	}
	if mqttCfg.Broker == "" || mqttCfg.Username == "" || mqttCfg.Password == "" {
		log.Fatal("MQTT_BROKER, MQTT_USERNAME, MQTT_PASSWORD must be set")
	}
	if _, err := mqtt.Connect(mqttCfg, hub); err != nil {
		log.Fatalf("❌ MQTT connection failed: %v", err)
	}
	// --- end MQTT setup ---

	h := setupHandlers(database, hub)
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
