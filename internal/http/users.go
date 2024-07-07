package http

import (
	"net/http"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/service/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		ctx.JSON(422, gin.H{
			"errors": []Invalid{
				{Field: "firstName, lastName, email and password required", Message: err.Error()},
			},
		})
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
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
		return
	}

	validUser, token, err := o.serv.LoginUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
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

func (o *UserController) GetUser(ctx *gin.Context) {
	var userIDPathParam UserIDPathParam
	if err := ctx.ShouldBindUri(&userIDPathParam); err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
		return
	}

	user, err := o.serv.GetUser(ctx, uuid.MustParse(userIDPathParam.UserID))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, BadUserRes{
			Status:     "Bad request",
			Message:    "failed to get user",
			StatusCode: 401,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get user successful",
		"data":    *GetUserRes(user),
	})
}
