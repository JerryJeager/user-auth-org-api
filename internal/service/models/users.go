package models

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
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
type GetUserRes struct {
	ID        uuid.UUID `json:"userId" gorm:"primary_key;type:uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique"`
	Phone     string    `json:"phone"`
}
type CreateUserReq struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone"`
}
type LoginUserReq struct {
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password" binding:"required"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
