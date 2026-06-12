package services

import (
	"kielio-gin-angular-app/backend/internal/models"
	"kielio-gin-angular-app/backend/internal/repositories"
)

type LanguageService struct {
	Repo repositories.LanguageRepository
}

func (s LanguageService) GetAllUsers() ([]models.Language, error) {
	return s.Repo.GetAll()
}

func (s LanguageService) GetByName(name string) (*models.Language, error) {
	return s.Repo.GetByName(name)
}
