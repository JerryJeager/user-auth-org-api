package organisations

import (
	"context"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/google/uuid"
)

type OrgSv interface {
	CreateOrganisation(ctx context.Context, org *models.Organisation, userID uuid.UUID) (*models.Organisation, error)
	CreateOrgMember(ctx context.Context, orgID, userID uuid.UUID) error
	GetOrganisation(ctx context.Context, orgID uuid.UUID) (*models.Organisation, error)
}

type OrgServ struct {
	repo OrgStore
}

func NewOrgService(repo OrgStore) *OrgServ {
	return &OrgServ{repo: repo}
}

func (o *OrgServ) CreateOrganisation(ctx context.Context, org *models.Organisation, userID uuid.UUID) (*models.Organisation, error) {
	id := uuid.New()
	org.ID = id
	org.UserID = userID
	var member = models.Member{
		UserID:         userID,
		OrganisationID: id,
	}

	err := o.repo.CreateOrganisation(ctx, org, &member)
	if err != nil {
		return &models.Organisation{}, err
	}

	return org, nil
}

func (o *OrgServ) CreateOrgMember(ctx context.Context, orgID, userID uuid.UUID) error {
	var member = models.Member{
		UserID:         userID,
		OrganisationID: orgID,
	}
	return o.repo.CreateOrgMember(ctx, &member)
}

func (o *OrgServ) GetOrganisation(ctx context.Context, orgID uuid.UUID) (*models.Organisation, error){
	return o.repo.GetOrganisation(ctx, orgID)
}