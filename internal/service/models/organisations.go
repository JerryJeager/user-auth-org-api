package models

import (
	"time"

	"github.com/google/uuid"
)

type Organisations struct {
	ID          uuid.UUID  `json:"orgId" gorm:"primary_key;type:uuid"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"user_id" gorm:"primary_key;type:uuid"`
}
