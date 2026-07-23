package router

import (
	"entrywatchserver/internal/handlers"
	"entrywatchserver/internal/ws"
	"net/http"
)

type Handlers struct {
	User  *handlers.UserHandler
	Athan *handlers.AttendanceHandler
	Hub   *ws.Hub
}

func New(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /users", h.User.GetAll)
	mux.HandleFunc("POST /users/new", h.User.AddUser)
	mux.HandleFunc("GET /records", h.Athan.GetAll)
	mux.HandleFunc("POST /records/new", h.Athan.RecordScan)

	mux.HandleFunc("GET /ws", h.Hub.ServeWS)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
