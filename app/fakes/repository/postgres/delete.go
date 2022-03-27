package fakes

import (
	"gokes/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *FakesRepository) Delete(fakesM *models.Fakes, tx *gorm.DB) error {
	var db = r.DB
	if tx != nil {
		db = tx
	}
	err := db.Delete(fakesM).Error
	if err != nil {
		log.Println("error-deleted-update:", err)

		return err
	}
	return nil
}
