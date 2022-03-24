package user

import (
	"gokes/app/models"
	"gokes/pkg/utils"
	"log"
)

func (r *UserRepository) Count(config utils.PaginationConfig) (int64, error) {
	var count int64

	err := r.DB.
		Model(&models.User{}).
		Scopes(config.Scopes()...).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-fakes:", err)
		return 0, err
	}

	return count, nil
}
