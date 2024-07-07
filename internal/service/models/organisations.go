package models

import (
	"time"

	"github.com/google/uuid"
)

type Organisation struct {
	ID          uuid.UUID  `json:"orgId" gorm:"primary_key;type:uuid"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"-" gorm:"type:uuid"`
}

type Organisations []Organisation

type OrganisationsRes struct {
	Organisation Organisations `json:"organisations"`
}

type Member struct {
	UserID         uuid.UUID `json:"user_id" gorm:"primary_key;type:uuid"`
	OrganisationID uuid.UUID `json:"organisation_id" gorm:"primary_key;type:uuid"`
}

type Members []Member

type CreateMemberReq struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
}
