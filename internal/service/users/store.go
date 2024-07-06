package users

import (
	"context"

	"github.com/JerryJeager/user-auth-org-api/config"
	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	result := config.Session.Create(user).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := config.Session.First(&user, "email = ?", email).WithContext(ctx)
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}
