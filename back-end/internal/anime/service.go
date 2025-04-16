package anime

import (
	"context"
	"errors"
)

type AnimeService interface {
	ProcessAttempt(ctx context.Context, currentCharacter any, name string, animeName string) (any, error)
	GenerateRandomCharacter(ctx context.Context, anime_name string) (any, error)
}

type animeServiceStruct struct {
	repo AnimeRepository
}

func NewAnimeService(repo AnimeRepository) AnimeService {
	return &animeServiceStruct{repo: repo}
}

func (s *animeServiceStruct) GenerateRandomCharacter(ctx context.Context, animeName string) (any, error) {
	if animeName == "" {
		return nil, errors.New("anime name cannot be empty")
	}

	character, err := s.repo.GetRandomCharacter(ctx, animeName)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (s *animeServiceStruct) ProcessAttempt(ctx context.Context, currentCharacter any, attemptedName string, animeName string) (any, error) {
	guessedCharacter, err := s.repo.GetAnimeCharacter(ctx, attemptedName, animeName)
	if err != nil {
		return nil, err
	}

	switch animeName {
	case "Naruto":
		selectedNaruto, ok1 := currentCharacter.(Naruto)
		guessedNaruto, ok2 := guessedCharacter.(Naruto)

		if !ok1 || !ok2 {
			return nil, errors.New("type assertion failed for Naruto")
		}

		result := NarutoComparisonResult{
			Name:             guessedNaruto.Name,
			Species:          compareStrings(selectedNaruto.Species, guessedNaruto.Species),
			PlaceOrigin:      compareStrings(selectedNaruto.PlaceOrigin, guessedNaruto.PlaceOrigin),
			IntroArc:         compareStrings(selectedNaruto.IntroArc, guessedNaruto.IntroArc),
			Affiliation:      compareStrings(selectedNaruto.Affiliation, guessedNaruto.Affiliation),
			ChakraType:       compareStrings(selectedNaruto.ChakraType, guessedNaruto.ChakraType),
			KekkeiGenkai:     compareStrings(selectedNaruto.KekkeiGenkai, guessedNaruto.KekkeiGenkai),
			JutsuAffinity:    compareStrings(selectedNaruto.JutsuAffinity, guessedNaruto.JutsuAffinity),
			SpecialAttribute: compareStrings(selectedNaruto.SpecialAttribute, guessedNaruto.SpecialAttribute),
		}
		return result, nil

	case "One Piece":
		selectedOnePiece, ok1 := currentCharacter.(OnePiece)
		guessedOnePiece, ok2 := guessedCharacter.(OnePiece)
		if !ok1 || !ok2 {
			return nil, errors.New("type assertion failed for One Piece")
		}
		result := OnePieceComparisonResult{
			Name:        selectedOnePiece.Name,
			Species:     compareStrings(selectedOnePiece.Species, guessedOnePiece.Species),
			PlaceOrigin: compareStrings(selectedOnePiece.PlaceOrigin, guessedOnePiece.PlaceOrigin),
			IntroArc:    compareStrings(selectedOnePiece.IntroArc, guessedOnePiece.IntroArc),
			Affiliation: compareStrings(selectedOnePiece.Affiliation, guessedOnePiece.Affiliation),
			Bounty:      compareInts(selectedOnePiece.Bounty, guessedOnePiece.Bounty),
			Haki:        compareStrings(selectedOnePiece.Haki, guessedOnePiece.Haki),
			DevilFruit:  compareStrings(selectedOnePiece.DevilFruit, guessedOnePiece.DevilFruit),
			Height:      compareInts(selectedOnePiece.Height, guessedOnePiece.Height),
		}
		return result, nil
	}
	return nil, nil
}
