package user

import (
	"gokes/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *UserRepository) Insert(userM *models.User, tx *gorm.DB) error {

	var db = r.DB
	if tx != nil {
		db = tx
	}
	err := db.Create(userM).Error
	if err != nil {
		log.Println("error-insert-user:", err)
		return err
	}
	return nil
}
