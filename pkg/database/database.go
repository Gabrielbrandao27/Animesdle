package database

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectToDatabase() (*sql.DB, error) {
	pass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/animesdb?parseTime=true", pass)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

func RunMigrations(db *sql.DB) error {
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

	if _, err := db.Exec(naruto_table); err != nil {
		return err
	}

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

	if _, err := db.Exec(onepiece_table); err != nil {
		return err
	}

	insert_naruto := `
				INSERT INTO characters_naruto (img_ref, name, species, place_origin, intro_arc, affiliation, chakra_type, kekkei_genkai, jutsu_affinity, special_attribute)
				VALUES
					("example.jpg", "Sasuke Uchiha", "Human", "Konohagakure", "Prologue", "Team 7, Taka, Akatsuki", "Fire Release", "Sharingan", "Ninjutsu, Genjutsu, Taijutsu", "Rinnegan")
				ON DUPLICATE KEY UPDATE name = VALUES(name);`
	if _, err := db.Exec(insert_naruto); err != nil {
		return err
	}

	insert_onepiece := `
				INSERT INTO characters_onepiece (img_ref, name, species, place_origin, intro_arc, affiliation, bounty, haki, devil_fruit, height)
				VALUES 
					("example.jpg", "Monkey D. Luffy", "Human", "Foosha Village", "Romance Dawn", "Straw Hat Pirates", 1500000000, "Observation, Armament, Conquerour", "Hito-Hito no mi: model Nika", 174),
					("example.jpg", "Roronoa Zoro", "Human", "Shimotsuki Village", "Romance Dawn", "Straw Hat Pirates", 1111000000, "Observation, Armament", "None", 181)
				ON DUPLICATE KEY UPDATE name = VALUES(name);`
	if _, err := db.Exec(insert_onepiece); err != nil {
		return err
	}

	return nil

}
