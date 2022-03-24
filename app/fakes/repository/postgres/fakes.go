package fakes

import "gorm.io/gorm"

// user struct for queries
type FakesRepository struct {
	DB *gorm.DB
}
