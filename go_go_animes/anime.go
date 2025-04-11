package go_go_animes

import (
	"errors"
	"fmt"
	"math/rand"
)

func AnimeCharacter(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was given, print it out.
	character := fmt.Sprintf("The character's name is %s", name)

	return character, nil
}

func RandomAnimeCharacter() (string, error) {
	// Generate a random anime character
	character := randomCharacter()

	// If no character was generated, return an error with a message.
	if character == "" {
		return "", errors.New("no character generated")
	}

	// If a character was generated, print it out.
	message := fmt.Sprintf("The character's name is %s", character)

	return message, nil
}

func randomCharacter() string {
	// List of anime characters
	characters := []string{
		"Naruto Uzumaki",
		"Sasuke Uchiha",
		"Sakura Haruno",
		"Monkey D. Luffy",
		"Zoro Roronoa",
		"Nami",
	}

	// Generate a random index
	index := rand.Intn(len(characters))

	return characters[index]
}
