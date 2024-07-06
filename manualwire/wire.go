package manualwire

import (
	"github.com/JerryJeager/user-auth-org-api/config"
	"github.com/JerryJeager/user-auth-org-api/internal/http"
	"github.com/JerryJeager/user-auth-org-api/internal/service/organisations"
	"github.com/JerryJeager/user-auth-org-api/internal/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}
func GetOrgRepository() *organisations.OrgRepo {
	repo := config.GetSession()
	return organisations.NewOrgRepo(repo)
}

func GetOrgService(repo organisations.OrgStore) *organisations.OrgServ {
	return organisations.NewOrgService(repo)
}

func GetOrgController() *http.OrgController {
	repo := GetOrgRepository()
	service := GetOrgService(repo)
	return http.NewOrgController(service)
}
