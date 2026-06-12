package repositories

import (
	"context"

	"kielio-gin-angular-app/backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type LanguageRepository struct {
	DB *pgx.Conn
}

func (r LanguageRepository) GetAll() ([]models.Language, error) {

	rows, err := r.DB.Query(
		context.Background(),
		"SELECT oid, name FROM language",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var languages []models.Language

	for rows.Next() {
		var l models.Language

		rows.Scan(
			&l.Oid,
			&l.Name,
		)

		languages = append(languages, l)
	}

	return languages, nil
}

func (r LanguageRepository) GetByName(name string) (*models.Language, error) {

	var language models.Language

	err := r.DB.QueryRow(
		context.Background(),
		`SELECT oid, name FROM language WHERE name = $1`,
		name,
	).Scan(
		&language.Oid,
		&language.Name,
	)

	if err != nil {
		return nil, err
	}

	return &language, nil
}
