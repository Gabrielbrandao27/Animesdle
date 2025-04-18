package anime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type AnimeHandler struct {
	service AnimeService
}

type AdminAnimeHandler struct {
	admin AdminService
}

type AttemptRequest struct {
	AttemptedName    string          `json:"name"`
	Anime            string          `json:"anime"`
	CurrentCharacter json.RawMessage `json:"currentCharacter"`
}

func NewAnimeHandler(service AnimeService) *AnimeHandler {
	return &AnimeHandler{service: service}
}

func NewAdminAnimeHandler(admin AdminService) *AdminAnimeHandler {
	return &AdminAnimeHandler{admin: admin}
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
	fmt.Println("DEBUG Anime:", req.Anime)
	fmt.Println("DEBUG Attempted Name:", req.AttemptedName)
	fmt.Println("DEBUG CurrentCharacter JSON:", string(req.CurrentCharacter))

	var currentCharacter any
	switch req.Anime {
	case "Naruto":
		fmt.Println("DEBUG CurrentCharacter JSON:", string(req.CurrentCharacter))
		var character Naruto
		if err := json.Unmarshal(req.CurrentCharacter, &character); err != nil {
			fmt.Println("Unmarshal error (Naruto):", err)
			http.Error(w, "invalid current character for Naruto", http.StatusBadRequest)
			return
		}
		fmt.Printf("Recebido do frontend:\nAnime: %s\nTentativa: %s\nCharacter: %s\n\n", req.Anime, req.AttemptedName, string(req.CurrentCharacter))
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

func (a *AdminAnimeHandler) DeleteRowsHandler(w http.ResponseWriter, r *http.Request) {
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
		err := a.admin.DeleteRows(r.Context(), anime, nil)
		if err != nil {
			http.Error(w, "error deleting rows: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := a.admin.DeleteRows(r.Context(), anime, id)
		if err != nil {
			http.Error(w, "error deleting rows: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Rows deleted successfully"))
}

func (a *AdminAnimeHandler) DropTableHandler(w http.ResponseWriter, r *http.Request) {
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

	err := a.admin.DropTable(r.Context(), anime)
	if err != nil {
		http.Error(w, "error dropping table: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Table dropped successfully"))
}

func (a *AdminAnimeHandler) AlterColumnSizeHandler(w http.ResponseWriter, r *http.Request) {
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

	column := r.URL.Query().Get("column")
	if column == "" {
		http.Error(w, "column parameter is required", http.StatusBadRequest)
		return
	}

	newSize := r.URL.Query().Get("newSize")
	if newSize == "" {
		http.Error(w, "newSize parameter is required", http.StatusBadRequest)
		return
	}

	err := a.admin.AlterColumnSize(r.Context(), anime, column, newSize)
	if err != nil {
		http.Error(w, "error altering column size: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Column size altered successfully"))
}
