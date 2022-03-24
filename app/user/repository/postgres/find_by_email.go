package user

import (
	"gokes/app/models"
)

// get user from email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	result := models.User{}

	err := r.DB.Where("email = ?", email).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
