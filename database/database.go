package database

import (
	"database/sql"
	"errors"
)

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:24H@R#gsd90ASF0@(127.0.0.1:3306)/animesdb?parseTime=true")
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

func CreateAnimeCharactersTable(db *sql.DB) {
	naruto_table := `
		CREATE TABLE IF NOT EXISTS characters_naruto (
			id INT AUTO_INCREMENT PRIMARY KEY,
			img_ref VARCHAR(255) NOT NULL,
			name VARCHAR(50) NOT NULL,
			species VARCHAR(50) NOT NULL,
			place_origin VARCHAR(50) NOT NULL,
			intro_arc VARCHAR(50) NOT NULL,
			affiliation VARCHAR(50) NOT NULL,
			chakra_type VARCHAR(50) NOT NULL,
			kekkei_genkai VARCHAR(50) NOT NULL,
			jutsu_affinity VARCHAR(50) NOT NULL,
			special_attribute VARCHAR(50) NOT NULL
		);`

	onepiece_table := `
		CREATE TABLE IF NOT EXISTS characters_onepiece (
			id INT AUTO_INCREMENT PRIMARY KEY,
			img_ref VARCHAR(255) NOT NULL,
			name VARCHAR(50) NOT NULL,
			species VARCHAR(50) NOT NULL,
			place_origin VARCHAR(50) NOT NULL,
			intro_arc VARCHAR(50) NOT NULL,
			affiliation VARCHAR(50) NOT NULL,
			bounty BIGINT UNSIGNED NOT NULL,
			haki VARCHAR(50) NOT NULL,
			devil_fruit VARCHAR(50) NOT NULL,
			height INT NOT NULL
		);`

	db.Exec(naruto_table)
	db.Exec(onepiece_table)
}

func InsertAnimeCharacter(db *sql.DB, payload map[string]interface{}, anime string) error {
	switch anime {
	case "Naruto":
		query := `
				INSERT INTO characters_naruto (img_ref, name, species, place_origin, intro_arc, affiliation, chakra_type, kekkei_genkai, jutsu_affinity, special_attribute)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		db.Exec(query,
			payload["img_ref"],
			payload["name"],
			payload["species"],
			payload["place_origin"],
			payload["intro_arc"],
			payload["affiliation"],
			payload["chakra_type"],
			payload["kekkei_genkai"],
			payload["jutsu_affinity"],
			payload["special_attribute"])

	case "One Piece":
		query := `
				INSERT INTO characters_onepiece (img_ref, name, species, place_origin, intro_arc, affiliation, bounty, haki, devil_fruit, height)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		db.Exec(query,
			payload["img_ref"],
			payload["name"],
			payload["species"],
			payload["place_origin"],
			payload["intro_arc"],
			payload["affiliation"],
			payload["bounty"],
			payload["haki"],
			payload["devil_fruit"],
			payload["height"])
	default:
		return errors.New("invalid anime type")
	}

	return nil
}

func GetAnimeCharacter(db *sql.DB, name string, anime string) (map[string]any, error) {
	switch anime {
	case "Naruto":
		var (
			id                                                                                                                    int
			img_ref, species, place_origin, intro_arc, affiliation, chakra_type, kekkei_genkai, jutsu_affinity, special_attribute string
		)
		query := "SELECT * FROM characters_naruto WHERE name = ?"
		err := db.QueryRow(query, name).Scan(&id, &img_ref, &name, &species, &place_origin, &intro_arc, &affiliation, &chakra_type, &kekkei_genkai, &jutsu_affinity, &special_attribute)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"id":                id,
			"img_ref":           img_ref,
			"name":              name,
			"species":           species,
			"place_origin":      place_origin,
			"intro_arc":         intro_arc,
			"affiliation":       affiliation,
			"chakra_type":       chakra_type,
			"kekkei_genkai":     kekkei_genkai,
			"jutsu_affinity":    jutsu_affinity,
			"special_attribute": special_attribute,
		}, nil

	case "One Piece":
		var (
			id                                                                        int
			img_ref, species, place_origin, intro_arc, affiliation, haki, devil_fruit string
			bounty, height                                                            int
		)
		query := "SELECT * FROM characters_onepiece WHERE name = ?"
		err := db.QueryRow(query, name).Scan(&id, &img_ref, &name, &species, &place_origin, &intro_arc, &affiliation, &bounty, &haki, &devil_fruit, &height)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"id":           id,
			"img_ref":      img_ref,
			"name":         name,
			"species":      species,
			"place_origin": place_origin,
			"intro_arc":    intro_arc,
			"affiliation":  affiliation,
			"bounty":       bounty,
			"haki":         haki,
			"devil_fruit":  devil_fruit,
			"height":       height,
		}, nil
	default:
		return nil, errors.New("invalid anime type")
	}
}
