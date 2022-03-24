package fakes

import (
	"gokes/app/models"
	"log"

	"gorm.io/gorm"
)

func (r *FakesRepository) Insert(fakesM *models.Fakes, tx *gorm.DB) error {

	var db = r.DB
	if tx != nil {
		db = tx
	}
	err := db.Create(fakesM).Error
	if err != nil {
		log.Println("error-insert-fakes:", err)
		return err
	}
	return nil
}
