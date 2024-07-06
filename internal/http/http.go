package http

import "github.com/JerryJeager/user-auth-org-api/internal/service/models"

type BadCreateUserRes struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int32  `json:"statusCode"`
}
type GoodCreateUserRes struct {
	AccessToken string `json:"accessToken"`
	User models.GetUserRes `json:"user"`
}

var ErrorCreatingUser = BadCreateUserRes{
	Status:     "Bad request",
	Message:    "Registration unsuccessful",
	StatusCode: 400,
}

func GetUserRes(user *models.User) *models.GetUserRes{
	return &models.GetUserRes{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Phone: user.Phone,
	}
}