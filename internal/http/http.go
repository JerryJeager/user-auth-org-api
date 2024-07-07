package http

import "github.com/JerryJeager/user-auth-org-api/internal/service/models"

type UserIDPathParam struct {
	UserID string `uri:"id" binding:"required,uuid_rfc4122"`
}
type OrgIDPathParam struct {
	OrgID string `uri:"orgId" binding:"required,uuid_rfc4122"`
}

type BadUserRes struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int32  `json:"statusCode"`
}
type GoodCreateUserRes struct {
	AccessToken string            `json:"accessToken"`
	User        models.GetUserRes `json:"user"`
}

type AllOrgsRes struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    models.OrganisationsRes `json:"data"`
}

var ErrorCreatingUser = BadUserRes{
	Status:     "Bad request",
	Message:    "Registration unsuccessful",
	StatusCode: 400,
}
var ErrorAuthUser = BadUserRes{
	Status:     "Bad request",
	Message:    "Authentication failed",
	StatusCode: 401,
}

type Invalid struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func GetUserRes(user *models.User) *models.GetUserRes {
	return &models.GetUserRes{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}
