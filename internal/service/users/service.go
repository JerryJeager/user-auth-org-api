package users

import (
	"context"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *models.CreateUserReq) (*models.User, string, error)
	LoginUser(ctx context.Context, user *models.LoginUserReq) (*models.User, string, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetYourUser(ctx context.Context, cUserId, userID uuid.UUID) (*models.User, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) CreateUser(ctx context.Context, user *models.CreateUserReq) (*models.User, string, error) {
	id := uuid.New()
	orgID := uuid.New()
	newUser := models.User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}
	newOrg := models.Organisation{
		ID: orgID,
		Name: utils.OrgName(user.FirstName),
		Description: "",
		UserID: id,
	}
	newMem := models.Member{
		UserID: id,
		OrganisationID: orgID,
	}

	if err := newUser.HashPassword(); err != nil {
		return &models.User{}, "", err
	}

	err := o.repo.CreateUser(ctx, &newUser, &newOrg, &newMem)
	if err != nil {
		return &models.User{}, "", err
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return &newUser, "", err
	}

	return &newUser, token, nil
}

func (o *UserServ) LoginUser(ctx context.Context, user *models.LoginUserReq) (*models.User, string, error) {
	validUser, err := o.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return &models.User{}, "", err
	}

	token, err := utils.GenerateToken(validUser.ID)
	if err != nil {
		return &models.User{}, "", err
	}

	return validUser, token, nil
}

func (o *UserServ) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return o.repo.GetUser(ctx, userID)
}

func (o *UserServ) GetYourUser(ctx context.Context, cUserId, userID uuid.UUID) (*models.User, error){
	return o.repo.GetYourUser(ctx, cUserId, userID)
}