package organisations

import "gorm.io/gorm"


type OrgStore interface {
	
}

type OrgRepo struct {
	client *gorm.DB
}

func NewOrgRepo(client *gorm.DB) *OrgRepo {
	return &OrgRepo{client: client}
}
