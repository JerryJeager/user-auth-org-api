package users

import (
	"context"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.CreateUserReq) (*models.User, string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) CreateUser(ctx context.Context, user *models.CreateUserReq) (*models.User, string, error) {
	id := uuid.New()
	newUser := models.User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}

	if err := newUser.HashPassword(); err != nil {
		return &models.User{}, "", err
	}

	err := o.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return &models.User{}, "", err
	}
	
	token, err :=  utils.GenerateToken(id)
	if err != nil{
		return &newUser, "", err 
	}

	return &newUser, token, nil
}
