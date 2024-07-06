package organisations

type OrgSv interface {
	
}

type OrgServ struct {
	repo OrgStore
}

func NewOrgService(repo OrgStore) *OrgServ {
	return &OrgServ{repo: repo}
}