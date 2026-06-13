package repositories

import (
	"context"

	"kielio-gin-angular-app/backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type DeutschAdjektivRepository struct {
	DB *pgx.Conn
}

func (r DeutschAdjektivRepository) GetAll() ([]models.DeutschAdjektiv, error) {

	rows, err := r.DB.Query(
		context.Background(),
		"SELECT * FROM deutsch_adjektiv",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return TransformToObject(rows), nil
}

func (r DeutschAdjektivRepository) GetByType(selectedType string) ([]models.DeutschAdjektiv, error) {

	rows, err := r.DB.Query(
		context.Background(),
		"SELECT * FROM deutsch_adjektiv WHERE type = $1",
		selectedType,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return TransformToObject(rows), nil
}

func (r DeutschAdjektivRepository) GetAllType() ([]string, error) {

	rows, err := r.DB.Query(
		context.Background(),
		`SELECT DISTINCT type FROM deutsch_adjektiv`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var types []string

	for rows.Next() {
		var articleType string
		err := rows.Scan(&articleType)

		if err != nil {
			return nil, err
		}

		types = append(types, articleType)
	}

	return types, nil
}

func TransformToObject(rows pgx.Rows) []models.DeutschAdjektiv {

	var deutschAdjektivs []models.DeutschAdjektiv

	for rows.Next() {
		var adj models.DeutschAdjektiv

		rows.Scan(
			&adj.Oid,
			&adj.Type,
			&adj.Gender,
			&adj.Case,
			&adj.ArticleEnding,
			&adj.AdjectiveEnding,
		)

		deutschAdjektivs = append(deutschAdjektivs, adj)
	}

	return deutschAdjektivs
}
