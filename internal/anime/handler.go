package anime

import (
	"context"
	"encoding/json"
	"net/http"
)

type AnimeHandler struct {
	service AnimeService
}

func NewAnimeHandler(service AnimeService) *AnimeHandler {
	return &AnimeHandler{service: service}
}

func (h *AnimeHandler) StartGameHandler(w http.ResponseWriter, r *http.Request) {
	anime := r.URL.Query().Get("anime")
	if anime == "" {
		http.Error(w, "anime parameter is required", http.StatusBadRequest)
		return
	}

	character, err := h.service.GenerateRandomCharacter(context.Background(), anime)
	if err != nil {
		http.Error(w, "error generating character: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}

func AttemptHandler(w http.ResponseWriter, r *http.Request) {
	// lógica da tentativa do usuário
}
