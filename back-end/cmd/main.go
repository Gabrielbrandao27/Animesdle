package main

import (
	"net/http"

	"github.com/Gabrielbrandao27/Animesdle/back-end/internal/anime"
	"github.com/Gabrielbrandao27/Animesdle/back-end/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	animeRepo := anime.NewAnimeRepository(db)
	animeService := anime.NewAnimeService(animeRepo)
	adminService := anime.NewAdminService(db)

	animeServiceHandler := anime.NewAnimeHandler(animeService)
	adminServiceHandler := anime.NewAdminAnimeHandler(adminService)

	http.HandleFunc("/start-game", animeServiceHandler.StartGameHandler)
	http.HandleFunc("/attempt", animeServiceHandler.AttemptHandler)
	http.HandleFunc("/admin/delete-rows", adminServiceHandler.DeleteRowsHandler)
	http.HandleFunc("/admin/drop-table", adminServiceHandler.DropTableHandler)
	http.HandleFunc("/admin/alter-column-size", adminServiceHandler.AlterColumnSizeHandler)
	http.ListenAndServe(":8080", nil)

}
