package organisations

import (
	"context"

	"github.com/JerryJeager/user-auth-org-api/config"
	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"gorm.io/gorm"
)

type OrgStore interface {
	CreateOrganisation(ctx context.Context, org *models.Organisation, mem *models.Member) error
	CreateOrgMember(ctx context.Context, member *models.Member) error
}

type OrgRepo struct {
	client *gorm.DB
}

func NewOrgRepo(client *gorm.DB) *OrgRepo {
	return &OrgRepo{client: client}
}

func (o *OrgRepo) CreateOrganisation(ctx context.Context, org *models.Organisation, mem *models.Member) error {
	err := config.Session.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(org).WithContext(ctx); result.Error != nil {
			return result.Error
		}
		if result := tx.Create(mem).WithContext(ctx); result.Error != nil {
			return result.Error
		}
		return nil
	})

	return err
}

func (o *OrgRepo) CreateOrgMember(ctx context.Context, member *models.Member) error {
	result := config.Session.Create(member).WithContext(ctx)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
