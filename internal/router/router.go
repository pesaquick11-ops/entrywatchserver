package router

import (
	"entrywatchserver/internal/handlers"
	"net/http"
)

type Handlers struct {
	User  *handlers.UserHandler
	Athan *handlers.AttendanceHandler
}

func New(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /users", h.User.GetAll)
	mux.HandleFunc("POST /users/new", h.User.AddUser)
	mux.HandleFunc("POST /record/new", h.Athan.RecordScan)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
