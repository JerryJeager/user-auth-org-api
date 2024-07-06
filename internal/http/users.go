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

func (o *UserController) LoginUser(ctx *gin.Context) {
	var user models.LoginUserReq
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorLoginUser)
		return
	}

	validUser, token, err := o.serv.LoginUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorLoginUser)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful",
		"data": GoodCreateUserRes{
			AccessToken: token,
			User:        *GetUserRes(validUser),
		},
	})
}
