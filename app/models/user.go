package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" groups:"user"`
	Username  string    `json:"username" groups:"user"`
	Password  string    `json:"password"`
	Email     string    `json:"email" groups:"user"`
	CreatedAt time.Time `json:"created_at" groups:"user"`
	UpdatedAt time.Time `json:"updated_at" groups:"user"`
}
