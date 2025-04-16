package anime

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
)

type AnimeRepository interface {
	GetAnimeCharacter(ctx context.Context, name string, anime string) (any, error)
	GetRandomCharacter(ctx context.Context, anime string) (any, error)
}

type mysqlAnimeRepository struct {
	db *sql.DB
}

func NewAnimeRepository(db *sql.DB) AnimeRepository {
	return &mysqlAnimeRepository{db: db}
}

func (r *mysqlAnimeRepository) GetAnimeCharacter(ctx context.Context, name string, anime_name string) (any, error) {
	switch anime_name {
	case "Naruto":
		query := "SELECT * FROM characters_naruto WHERE name = ?"
		var character Naruto
		err := r.db.QueryRowContext(ctx, query, name).Scan(
			&character.ID,
			&character.ImgRef,
			&character.Name,
			&character.Species,
			&character.PlaceOrigin,
			&character.IntroArc,
			&character.Affiliation,
			&character.ChakraType,
			&character.KekkeiGenkai,
			&character.JutsuAffinity,
			&character.SpecialAttribute,
		)
		if err != nil {
			return nil, err
		}
		return character, nil

	case "One Piece":
		query := "SELECT * FROM characters_onepiece WHERE name = ?"
		var character OnePiece
		err := r.db.QueryRowContext(ctx, query, name).Scan(
			&character.ID,
			&character.ImgRef,
			&character.Name,
			&character.Species,
			&character.PlaceOrigin,
			&character.IntroArc,
			&character.Affiliation,
			&character.Bounty,
			&character.Haki,
			&character.DevilFruit,
			&character.Height,
		)
		if err != nil {
			return nil, err
		}
		return character, nil
	default:
		return nil, errors.New("invalid anime type")
	}
}

func (r *mysqlAnimeRepository) GetRandomCharacter(ctx context.Context, anime string) (any, error) {
	var count int
	var queryCount, queryFetch string

	switch anime {
	case "Naruto":
		queryCount = "SELECT COUNT(*) FROM characters_naruto"
		queryFetch = "SELECT * FROM characters_naruto LIMIT 1 OFFSET ?"

		err := r.db.QueryRowContext(ctx, queryCount).Scan(&count)
		if err != nil {
			return nil, err
		}

		offset := rand.Intn(count)

		var character Naruto
		err = r.db.QueryRowContext(ctx, queryFetch, offset).Scan(
			&character.ID,
			&character.ImgRef,
			&character.Name,
			&character.Species,
			&character.PlaceOrigin,
			&character.IntroArc,
			&character.Affiliation,
			&character.ChakraType,
			&character.KekkeiGenkai,
			&character.JutsuAffinity,
			&character.SpecialAttribute,
		)
		if err != nil {
			return nil, err
		}

		return character, nil

	case "One Piece":
		queryCount = "SELECT COUNT(*) FROM characters_onepiece"
		queryFetch = "SELECT * FROM characters_onepiece LIMIT 1 OFFSET ?"

		err := r.db.QueryRowContext(ctx, queryCount).Scan(&count)
		if err != nil {
			return nil, err
		}

		offset := rand.Intn(count)

		var character OnePiece
		err = r.db.QueryRowContext(ctx, queryFetch, offset).Scan(
			&character.ID,
			&character.ImgRef,
			&character.Name,
			&character.Species,
			&character.PlaceOrigin,
			&character.IntroArc,
			&character.Affiliation,
			&character.Bounty,
			&character.Haki,
			&character.DevilFruit,
			&character.Height,
		)
		if err != nil {
			return nil, err
		}

		return character, nil
	}
	return nil, nil
}
