package http

import (
	"net/http"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/service/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}

func (o *UserController) CreateUser(ctx *gin.Context) {
	var user models.CreateUserReq

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorCreatingUser)
		return
	}

	newUser, token, err := o.serv.CreateUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorCreatingUser)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Registration successful",
		"data": GoodCreateUserRes{
			AccessToken: token,
			User:        *GetUserRes(newUser),
		},
	})
}
