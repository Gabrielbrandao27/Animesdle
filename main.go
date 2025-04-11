package main

import (
	"log"

	"github.com/Gabrielbrandao27/Animesdle/database"
	"github.com/Gabrielbrandao27/Animesdle/go_go_animes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	character, err := go_go_animes.AnimeCharacter("Ichigo")
	if err != nil {
		log.Fatal(err)
	}
	println(character)

	randomCharacter, err := go_go_animes.RandomAnimeCharacter()
	if err != nil {
		log.Fatal(err)
	}
	println(randomCharacter)
}

func init() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateAnimeCharactersTable(db)

	payload := map[string]any{
		"img_ref":           "https://example.com/image.jpg",
		"name":              "Naruto Uzumaki",
		"species":           "Human",
		"place_origin":      "Konohagakure",
		"intro_arc":         "Prologue",
		"affiliation":       "Team 7",
		"chakra_type":       "Wind",
		"kekkei_genkai":     "None",
		"jutsu_affinity":    "Shadow Clone Jutsu",
		"special_attribute": "Sage Mode",
	}

	err = database.InsertAnimeCharacter(db, payload, "Naruto")
	if err != nil {
		log.Fatal(err)
	}
}
