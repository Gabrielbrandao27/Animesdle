package anime

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type AnimeHandler struct {
	service AnimeService
}

type AnimeHandlerRepository struct {
	repository AnimeRepository
}

type AttemptRequest struct {
	AttemptedName    string          `json:"AttemptedName"`
	Anime            string          `json:"anime"`
	CurrentCharacter json.RawMessage `json:"currentCharacter"`
}

func NewAnimeHandler(service AnimeService) *AnimeHandler {
	return &AnimeHandler{service: service}
}

func NewAnimeRepositoryHandler(repository AnimeRepository) *AnimeHandlerRepository {
	return &AnimeHandlerRepository{repository: repository}
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

func (repo *AnimeHandlerRepository) DeleteRowsHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	expectedToken := os.Getenv("TOKEN")

	if token != expectedToken {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	anime := r.URL.Query().Get("anime")
	if anime == "" {
		http.Error(w, "anime parameter is required", http.StatusBadRequest)
		return
	}

	var id *int64
	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		parsedID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}
		id = &parsedID
	}

	if id == nil {
		err := repo.repository.DeleteRows(r.Context(), anime, nil)
		if err != nil {
			http.Error(w, "error deleting rows: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := repo.repository.DeleteRows(r.Context(), anime, id)
		if err != nil {
			http.Error(w, "error deleting rows: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Rows deleted successfully"))
}
