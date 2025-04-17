package main

import (
	"net/http"

	"github.com/Gabrielbrandao27/Animesdle/back-end/internal/anime"
	myhttp "github.com/Gabrielbrandao27/Animesdle/back-end/internal/http"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/start-game", animeServiceHandler.StartGameHandler)
	mux.HandleFunc("/attempt", animeServiceHandler.AttemptHandler)
	mux.HandleFunc("/admin/delete-rows", adminServiceHandler.DeleteRowsHandler)
	mux.HandleFunc("/admin/drop-table", adminServiceHandler.DropTableHandler)
	mux.HandleFunc("/admin/alter-column-size", adminServiceHandler.AlterColumnSizeHandler)

	handlerWithCORS := myhttp.CorsMiddleware(mux)

	http.ListenAndServe(":8080", handlerWithCORS)

}
