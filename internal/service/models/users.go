package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID  `json:"userId" gorm:"primary_key;type:uuid"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
}
