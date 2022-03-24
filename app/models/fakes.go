package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	FakesStatusRumahSakit = "rumah_sakit"
	FakesStatusPuskesmas  = "puskesmas"
	FakesStatusPosyandu   = "posyandu"
	FakesStatusKlinik     = "klinik"
)

type Fakes struct {
	ID          uuid.UUID      `json:"id" groups:"user,public"`
	Code        string         `json:"code" groups:"user,public"`
	Name        string         `json:"name" groups:"user,public"`
	Type        string         `json:"type" groups:"user,public"`
	Description string         `json:"description" groups:"user,public"`
	NakesCount  uint64         `json:"nakes_count" groups:"user,public"`
	CreatedBy   uuid.UUID      `json:"updated_by" groups:"user"`
	CreatedAt   time.Time      `json:"created_at" groups:"user,public"`
	UpdatedAt   time.Time      `json:"updated_at" groups:"user"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" groups:"user"`
}
