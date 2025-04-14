package anime

import (
	"context"
	"errors"
)

type AnimeService interface {
	ProcessAttempt(ctx context.Context, name string, anime_name string) (any, error)
	GenerateRandomCharacter(ctx context.Context, anime_name string) (any, error)
}

type animeServiceStruct struct {
	repo AnimeRepository
}

func NewAnimeService(repo AnimeRepository) AnimeService {
	return &animeServiceStruct{repo: repo}
}

func (s *animeServiceStruct) ProcessAttempt(ctx context.Context, name string, anime_name string) (any, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	character, err := s.repo.GetAnimeCharacter(ctx, name, anime_name)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (s *animeServiceStruct) GenerateRandomCharacter(ctx context.Context, anime_name string) (any, error) {
	if anime_name == "" {
		return nil, errors.New("anime name cannot be empty")
	}

	character, err := s.repo.GetRandomCharacter(ctx, anime_name)
	if err != nil {
		return nil, err
	}

	return character, nil
}
