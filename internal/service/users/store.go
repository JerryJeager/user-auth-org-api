package users

import (
	"context"
	"errors"

	"github.com/JerryJeager/user-auth-org-api/config"
	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User, org *models.Organisation, mem *models.Member) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetYourUser(ctx context.Context, cUserId, userID uuid.UUID) (*models.User, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUser(ctx context.Context, user *models.User, org *models.Organisation, mem *models.Member) error {
	err := config.Session.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(user).WithContext(ctx); result.Error != nil {
			return result.Error
		}
		if result := tx.Create(org).WithContext(ctx); result.Error != nil {
			return result.Error
		}
		if result := tx.Create(mem).WithContext(ctx); result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err
}

func (o *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := config.Session.First(&user, "email = ?", email).WithContext(ctx)
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}

// get's user record of logged in current user
func (o *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	result := config.Session.First(&user, "id = ?", userID).WithContext(ctx)
	if result.Error != nil {
		return &models.User{}, nil
	}
	return &user, nil
}

// get's user record of a user in the same organisation as the current logged in user who's making the request
func (o *UserRepo) GetYourUser(ctx context.Context, cUserId, userID uuid.UUID) (*models.User, error) {
	var currentUserOrgs models.Members
	var yourUserOrgs models.Members
	var yourUser models.User
	if result := config.Session.Find(&currentUserOrgs, "user_id = ?", cUserId); result.Error != nil {
		return &models.User{}, result.Error
	}
	if result := config.Session.Find(&yourUserOrgs, "user_id = ?", userID); result.Error != nil {
		return &models.User{}, result.Error
	}
	if sameOrgs := utils.IsInSameOrganisation(currentUserOrgs, yourUserOrgs); !sameOrgs {
		return &models.User{}, errors.New("users must be in the same organisation to see each other's record")
	}
	if result := config.Session.First(&yourUser, "id = ?", userID); result.Error != nil {
		return &models.User{}, result.Error
	}
	return &yourUser, nil
}
