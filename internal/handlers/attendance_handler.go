package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"entrywatchserver/internal/repository"
)

type AttendanceHandler struct {
	repo *repository.AttendanceRepository
}

func NewAttendanceHandler(attRepo *repository.AttendanceRepository) *AttendanceHandler {
	return &AttendanceHandler{repo: attRepo}
}

func (a *AttendanceHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	attendances, err := a.repo.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(attendances)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type ScanRequest struct {
	Username   string `json:"username"`
	Confidence string `json:"confidence"`
}

func (a *AttendanceHandler) RecordScan(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body ", http.StatusBadRequest)
		return
	}
	if req.Username == "" {
		http.Error(w, "Username is not provided", http.StatusBadRequest)
		return
	}
	if req.Confidence == "" {
		http.Error(w, "Confidence is not provided", http.StatusBadRequest)
		return
	}
	err = a.repo.RecordScan(r.Context(), req.Username, req.Confidence, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"timestamp": "added"})
}
