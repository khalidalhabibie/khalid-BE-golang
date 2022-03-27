package fakes

import (
	"gokes/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *FakesRepository) Update(fakesM *models.Fakes, tx *gorm.DB) error {
	var db = r.DB
	if tx != nil {
		db = tx
	}
	err := db.Save(fakesM).Error
	if err != nil {
		log.Println("error-insert-update:", err)
		return err
	}

	return nil
}
