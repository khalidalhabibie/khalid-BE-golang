package fakes

import (
	"gokes/app/models"
	"gokes/pkg/utils"
	"log"
)

func (r *FakesRepository) FindAll(config utils.PaginationConfig) ([]models.User, error) {
	results := []models.User{}

	err := r.DB.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error	
	if err != nil {
		log.Println("error-find-all-fakes:", err)
		return nil, err
	}

	return results, nil
}
