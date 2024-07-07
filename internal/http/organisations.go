package http

import (
	"net/http"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/JerryJeager/user-auth-org-api/internal/service/organisations"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrgController struct {
	serv organisations.OrgSv
}

func NewOrgController(serv organisations.OrgSv) *OrgController {
	return &OrgController{serv: serv}
}

func (o *OrgController) CreateOrganisation(ctx *gin.Context) {
	var org models.Organisation
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":     "Bad request",
			"statusCode": 400,
			"message":    "failed to create organisation",
		})
		return
	}

	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	var userID uuid.UUID
	if id.(string) == "" {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	userID, err := uuid.Parse(id.(string))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	organisation, err := o.serv.CreateOrganisation(ctx, &org, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":     "Bad request",
			"statusCode": 400,
			"message":    "failed to create organisation",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Organisation created successfully",
		"data":    *organisation,
	})

}

func (o *OrgController) CreateOrgMember(ctx *gin.Context) {
	var newMember models.CreateMemberReq
	if err := ctx.ShouldBindJSON(&newMember); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": "user id is required",
		})
		return
	}
	var orgID OrgIDPathParam
	if err := ctx.ShouldBindUri(&orgID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": "organisation id is required in path parameter",
		})
		return
	}

	err := o.serv.CreateOrgMember(ctx, uuid.MustParse(orgID.OrgID), newMember.UserID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": "failed to add user to organisation",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User added to organisation successfully",
	})
}

func (o *OrgController) GetOrganisation(ctx *gin.Context) {
	var orgIDPathParam OrgIDPathParam
	if err := ctx.ShouldBindUri(&orgIDPathParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": "organisation id in path parameter should be of type uuid",
		})
		return
	}

	organisation, err := o.serv.GetOrganisation(ctx, uuid.MustParse(orgIDPathParam.OrgID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "organisation not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get organisation successful",
		"data":    *organisation,
	})
}

func (o *OrgController) GetOrganisations(ctx *gin.Context) {
	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(id.(string))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	organisations, err := o.serv.GetOrganisations(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "organisation not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get all organisations successful",
		"data": models.OrganisationsRes{
			Organisation: *organisations,
		},
	})
}
