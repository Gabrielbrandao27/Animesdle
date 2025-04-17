package main

import (
	"log"

	"github.com/Gabrielbrandao27/Animesdle/back-end/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Erro nas migrations:", err)
	}

	log.Println("Migrations executadas com sucesso!")
}
