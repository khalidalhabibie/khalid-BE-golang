package user

import (
	"gokes/app/models"
)

// get user from email
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	result := models.User{}

	err := r.DB.Where("username = ?", username).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
