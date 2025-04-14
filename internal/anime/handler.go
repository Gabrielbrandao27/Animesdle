package anime

import (
	"context"
	"encoding/json"
	"net/http"
)

type AnimeHandler struct {
	service AnimeService
}

type AttemptRequest struct {
	AttemptedName    string          `json:"AttemptedName"`
	Anime            string          `json:"anime"`
	CurrentCharacter json.RawMessage `json:"currentCharacter"`
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

func (h *AnimeHandler) AttemptHandler(w http.ResponseWriter, r *http.Request) {
	var req AttemptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var currentCharacter any
	switch req.Anime {
	case "Naruto":
		var character Naruto
		if err := json.Unmarshal(req.CurrentCharacter, &character); err != nil {
			http.Error(w, "invalid current character for Naruto", http.StatusBadRequest)
			return
		}
		currentCharacter = character
	case "One Piece":
		var character OnePiece
		if err := json.Unmarshal(req.CurrentCharacter, &character); err != nil {
			http.Error(w, "invalid current character for One Piece", http.StatusBadRequest)
			return
		}
		currentCharacter = character
	}

	// Chama o serviço com os dados
	result, err := h.service.ProcessAttempt(r.Context(), currentCharacter, req.AttemptedName, req.Anime)
	if err != nil {
		http.Error(w, "error processing attempt: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
