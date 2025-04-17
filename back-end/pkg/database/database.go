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
			name VARCHAR(100) NOT NULL UNIQUE,
			species VARCHAR(100) NOT NULL,
			place_origin VARCHAR(100) NOT NULL,
			intro_arc VARCHAR(255) NOT NULL,
			affiliation VARCHAR(255) NOT NULL,
			chakra_type VARCHAR(255) NOT NULL,
			kekkei_genkai VARCHAR(255) NOT NULL,
			jutsu_affinity VARCHAR(255) NOT NULL,
			special_attribute VARCHAR(255) NOT NULL
		);`

	if _, err := db.Exec(naruto_table); err != nil {
		return err
	}

	onepiece_table := `
		CREATE TABLE IF NOT EXISTS characters_onepiece (
			id INT AUTO_INCREMENT PRIMARY KEY,
			img_ref VARCHAR(255) NOT NULL,
			name VARCHAR(100) NOT NULL UNIQUE,
			species VARCHAR(255) NOT NULL,
			place_origin VARCHAR(255) NOT NULL,
			intro_arc VARCHAR(100) NOT NULL,
			affiliation VARCHAR(255) NOT NULL,
			bounty BIGINT UNSIGNED NOT NULL,
			haki VARCHAR(100) NOT NULL,
			devil_fruit VARCHAR(255) NOT NULL,
			height INT NOT NULL
		);`

	if _, err := db.Exec(onepiece_table); err != nil {
		return err
	}

	insert_naruto := `
				INSERT INTO characters_naruto (img_ref, name, species, place_origin, intro_arc, affiliation, chakra_type, kekkei_genkai, jutsu_affinity, special_attribute)
				VALUES
					("example.jpg", "Naruto Uzumaki", "Human", "Konohagakure", "Prologue", "Team 7, Kage, Mount Myōboku", "Wind", "None", "Ninjutsu, Taijutsu", "Sage, Jinchūriki"),
					("example.jpg", "Sasuke Uchiha", "Human", "Konohagakure", "Prologue", "Team 7, Taka, Akatsuki", "Earth, Fire, Wind, Water, Lightning", "Sharingan", "Ninjutsu, Genjutsu, Taijutsu", "Eternal Mangekyō Sharingan, Rinnegan, Rinnesharingan"),
					("example.jpg", "Sakura Haruno", "Human", "Konohagakure", "Prologue", "Team 7", "Wind", "None", "Taijutsu, Genjutsu", "Creation Rebirth"),
					("example.jpg", "Kakashi Hatake", "Human", "Konohagakure", "Prologue", "Team Minato, Team 7, ANBU, Kage", "Lightning", "Sharingan", "Ninjutsu, Taijutsu, Genjutsu", "Mangekyō Sharingan"),
					("example.jpg", "Hashirama Senju", "Human", "Konohagakure", "Konoha Crush", "Kage", "Water, Earth", "Wood", "Ninjutsu, Taijutsu", "Sage"),
					("example.jpg", "Tobirama Senju", "Human", "Konohagakure", "Konoha Crush", "Kage", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Edo Tensei"),
					("example.jpg", "Hiruzen Sarutobi", "Human", "Konohagakure", "Prologue", "Kage", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "None"),
					("example.jpg", "Minato Namikaze", "Human", "Konohagakure", "Kakashi Gaiden", "Team Minato, Kage, Mount Myōboku", "Fire, Wind, Lightning", "None", "Ninjutsu, Taijutsu", "Sage, Jinchūriki"),
					("example.jpg", "Tsunade Senju", "Human", "Konohagakure", "Search For Tsunade", "Kage, Legendary Sannin", "Earth, Fire, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Creation Rebirth"),
					("example.jpg", "Jiraiya", "Human", "Konohagakure", "Chunnin Exams", "Legendary Sannin, Mount Myōboku", "Earth, Fire, Water, Wind", "None", "Ninjutsu, Taijutsu, Genjutsu", "Sage"),
					("example.jpg", "Madara Uchiha", "Human", "Konohagakure", "Tale of Jiraiya the Gallant", "Akatsuki", "Earth, Fire, Wind, Water, Lightning", "Sharingan, Wood, Storm", "Ninjutsu, Taijutsu, Genjutsu", "Sage, Eternal Mangekyō Sharingan, Rinnegan, Rinnesharingan, Jinchūriki"),
					("example.jpg", "Haku", "Human", "Kirigakure", "Land of Waves", "Mercenary", "Water", "Ice", "Ninjutsu, Taijutsu, Genjutsu", "None"),
					("example.jpg", "Zabuza Momochi", "Human", "Kirigakure", "Land of Waves", "Missing-nin, Mercenary, Seven Ninja Swordsmen of the Mist", "Water", "None", "Ninjutsu, Taijutsu, Genjutsu", "None"),
					("example.jpg", "Rock Lee", "Human", "Konohagakure", "Prologue", "Team Guy", "Earth, Fire, Wind, Water, Lightning", "None", "Taijutsu, Ninjutsu, Genjutsu", "None"),
					("example.jpg", "Neji Hyuga", "Human", "Konohagakure", "Prologue", "Team Guy", "Earth, Fire, Wind, Water, Lightning", "Byakugan", "Ninjutsu, Taijutsu, Genjutsu", "None"),
					("example.jpg", "Tenten", "Human", "Konohagakure", "Prologue", "Team Guy", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "None"),
					("example.jpg", "Shikamaru Nara", "Human", "Konohagakure", "Prologue", "Team 10", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Shadow Manipulation"),
					("example.jpg", "Ino Yamanaka", "Human", "Konohagakure", "Prologue", "Team 10", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Mind Transfer Jutsu"),
					("example.jpg", "Choji Akimichi", "Human", "Konohagakure", "Prologue", "Team 10", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Expansion Jutsu"),
					("example.jpg", "Kiba Inuzuka", "Human", "Konohagakure", "Prologue", "Team 8", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Beast Mimicry"),
					("example.jpg", "Hinata Hyuga", "Human", "Konohagakure", "Prologue", "Team 8", "Earth, Fire, Wind, Water, Lightning", "Byakugan", "Ninjutsu, Taijutsu, Genjutsu", "Gentle Fist"),
					("example.jpg", "Shino Aburame", "Human", "Konohagakure", "Prologue", "Team 8", "Earth, Fire, Wind, Water, Lightning", "None", "Ninjutsu, Taijutsu, Genjutsu", "Insect Manipulation")
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
