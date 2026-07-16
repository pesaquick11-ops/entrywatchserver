package main

import (
	"entrywatchserver/internal/db"
	"fmt"
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"hello world")
}

func main() {

	mongoClient,err := db.ConnectToMongo("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database("test")
	


	http.HandleFunc("/", rootHandler)

	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
