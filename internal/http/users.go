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
	var currentUserID uuid.UUID
	var user *models.User
	var getUserError error
	if err := ctx.ShouldBindUri(&userIDPathParam); err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
		return
	}
	id, found := ctx.Get("user_id")
	if !found {
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
		return
	}
	currentUserID, err := uuid.Parse(id.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorAuthUser)
		return
	}

	if currentUserID.String() == userIDPathParam.UserID {
		user, getUserError = o.serv.GetUser(ctx, uuid.MustParse(userIDPathParam.UserID))
	} else {
		user, getUserError = o.serv.GetYourUser(ctx, currentUserID, uuid.MustParse(userIDPathParam.UserID))
	}

	if getUserError != nil {
		ctx.JSON(http.StatusForbidden, BadUserRes{
			Status:     "Bad request",
			Message:    getUserError.Error(),
			StatusCode: http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get user successful",
		"data":    *GetUserRes(user),
	})
}
