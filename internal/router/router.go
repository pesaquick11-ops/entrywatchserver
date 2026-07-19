package router

import (
	"entrywatchserver/internal/handlers"
	"net/http"
)

type Handlers struct {
	User *handlers.UserHandler
}

func New(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /users", h.User.GetAll)
	mux.HandleFunc("POST /users/new", h.User.AddUser)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
