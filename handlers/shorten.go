package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yashK-2703/go-url-shortener/storage"
	"github.com/yashK-2703/go-url-shortener/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := utils.GenerateID()
	err := storage.SetURL(id, req.URL)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"short_url": "http://localhost:8080/r/" + id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
