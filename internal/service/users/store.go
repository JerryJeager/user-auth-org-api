package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/JerryJeager/user-auth-org-api/config"
	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
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

// get's user record of another user only if they are in the same organisation
func (o *UserRepo) GetYourUser(ctx context.Context, cUserId, userID uuid.UUID) (*models.User, error) {
	var yourUser models.User
	query := fmt.Sprintf(`
	select
	 	u.id, u.first_name, u.last_name, u.email, u.created_at, u.updated_at, u.deleted_at
	 	from users as u
    	inner join (
   			select y.user_id as your_user_id, y.organisation_id as your_user_org_id
			from members as y
    		inner join 
     			(select * from members where user_id = '%s') as c 
     			on y.organisation_id = c.organisation_id
    		where y.user_id = '%s'
		) 
		as y on y.your_user_id = u.id`, cUserId, userID) 

	result := config.Session.Raw(query).Scan(&yourUser) 
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	if result.RowsAffected < 1{
		return &models.User{}, errors.New("users must be in the same organisation to see each other's record")
	}
	return &yourUser, nil
}
