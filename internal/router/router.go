package router

import (
	"net/http"
	"entrywatchserver/internal/handlers"
)

func New(userHandler *handlers.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /users", userHandler.GetAll)
	// mux.HandleFunc("GET /products", productHandler.GetAll)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}