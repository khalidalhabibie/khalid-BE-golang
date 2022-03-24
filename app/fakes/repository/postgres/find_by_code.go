package fakes

import (
	"gokes/app/models"
)

// get user from email
func (r *FakesRepository) FindByCode(code string) (*models.Fakes, error) {
	result := models.Fakes{}

	err := r.DB.Where("code = ?", code).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
