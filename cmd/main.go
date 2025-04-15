package main

import (
	"net/http"

	"github.com/Gabrielbrandao27/Animesdle/internal/anime"
	"github.com/Gabrielbrandao27/Animesdle/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	animeRepo := anime.NewAnimeRepository(db)
	animeService := anime.NewAnimeService(animeRepo)

	animeRepositoryHandler := anime.NewAnimeRepositoryHandler(animeRepo)
	animeServiceHandler := anime.NewAnimeHandler(animeService)

	http.HandleFunc("/start-game", animeServiceHandler.StartGameHandler)
	http.HandleFunc("/attempt", animeServiceHandler.AttemptHandler)
	http.HandleFunc("/delete", animeRepositoryHandler.DeleteRowsHandler)
	http.ListenAndServe(":8080", nil)

}
