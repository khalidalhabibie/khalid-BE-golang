package fakes

import (
	"gokes/app/models"
	"log"
)

// get user from code
func (r *FakesRepository) FindByCode(code string) (*models.Fakes, error) {
	result := models.Fakes{}

	err := r.DB.Where("code = ?", code).First(&result).Error

	log.Println("error repos ", err)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
