package fakes

import (
	"gokes/app/models"

	"github.com/google/uuid"
)

// get user from email
func (r *FakesRepository) FindByID(id uuid.UUID) (*models.Fakes, error) {
	result := models.Fakes{}

	err := r.DB.Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
