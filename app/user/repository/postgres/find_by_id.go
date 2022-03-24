package user

import (
	"gokes/app/models"

	"github.com/google/uuid"
)

// get user from email
func (r *UserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	result := models.User{}

	err := r.DB.Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
