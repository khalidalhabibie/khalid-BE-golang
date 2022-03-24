package user

import "gorm.io/gorm"

// user struct for queries
type UserRepository struct {
	DB *gorm.DB
}
