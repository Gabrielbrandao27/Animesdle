package anime

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type AdminService interface {
	DeleteRows(ctx context.Context, anime string, id *int64) error
	DropTable(ctx context.Context, anime string) error
	AlterColumnSize(ctx context.Context, anime string, column string, newSize string) error
}

type mysqlAdminService struct {
	db *sql.DB
}

func NewAdminService(db *sql.DB) AdminService {
	return &mysqlAdminService{db: db}
}

func (a *mysqlAdminService) DeleteRows(ctx context.Context, anime string, id *int64) error {
	var query string

	switch anime {
	case "Naruto":
		if id != nil {
			query = "DELETE FROM characters_naruto WHERE id = ?"
			_, err := a.db.ExecContext(ctx, query, *id)
			return err
		}
		query = "DELETE FROM characters_naruto"
	case "One Piece":
		if id != nil {
			query = "DELETE FROM characters_onepiece WHERE id = ?"
			_, err := a.db.ExecContext(ctx, query, *id)
			return err
		}
		query = "DELETE FROM characters_onepiece"
	default:
		return errors.New("invalid anime type")
	}

	_, err := a.db.ExecContext(ctx, query)
	return err
}

func (a *mysqlAdminService) DropTable(ctx context.Context, anime string) error {
	var query string

	switch anime {
	case "Naruto":
		query = "DROP TABLE IF EXISTS characters_naruto"
	case "One Piece":
		query = "DROP TABLE IF EXISTS characters_onepiece"
	default:
		return errors.New("invalid anime type")
	}

	_, err := a.db.ExecContext(ctx, query)
	return err
}

func (r *mysqlAdminService) AlterColumnSize(ctx context.Context, anime string, column string, newSize string) error {
	var query string

	switch anime {
	case "Naruto":
		query = fmt.Sprintf("ALTER TABLE characters_naruto MODIFY %s VARCHAR(%s)", column, newSize)
	case "One Piece":
		query = fmt.Sprintf("ALTER TABLE characters_onepiece MODIFY %s VARCHAR(%s)", column, newSize)
	default:
		return errors.New("invalid anime type")
	}

	_, err := r.db.ExecContext(ctx, query)
	return err
}
