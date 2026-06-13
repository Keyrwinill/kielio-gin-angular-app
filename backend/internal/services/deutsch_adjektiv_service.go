package services

import (
	"kielio-gin-angular-app/backend/internal/dto"
	"kielio-gin-angular-app/backend/internal/models"
	"kielio-gin-angular-app/backend/internal/repositories"
)

type DeutschAdjektivService struct {
	Repo repositories.DeutschAdjektivRepository
}

func (s DeutschAdjektivService) GetByType(selectedType string) ([]dto.DeutschAdjektivResponse, error) {

	var modelsList []models.DeutschAdjektiv
	var err error

	if selectedType == "All" {
		modelsList, err = s.Repo.GetAll()
	} else {
		modelsList, err = s.Repo.GetByType(selectedType)
	}

	if err != nil {
		return nil, err
	}

	types, err := s.Repo.GetAllType()

	if err != nil {
		return nil, err
	}

	responseList := make([]dto.DeutschAdjektivResponse, 0, len(modelsList))

	for _, item := range modelsList {
		responseList = append(responseList, ToDeutschAdjektivResponse(item, types))
	}

	return responseList, nil
}

func ToDeutschAdjektivResponse(
	item models.DeutschAdjektiv,
	types []string,
) dto.DeutschAdjektivResponse {

	return dto.DeutschAdjektivResponse{
		Oid:             item.Oid.String(),
		Type:            item.Type,
		Gender:          item.Gender,
		Case:            item.Case,
		ArticleEnding:   item.ArticleEnding,
		AdjectiveEnding: item.AdjectiveEnding,
		ArticleTypeList: types,
	}
}