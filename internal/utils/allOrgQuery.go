package utils

import (
	"fmt"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
)

// generates where condition clause for getting all organisations for a user from an array containing all organisations id

func AllOrgQuery(members models.Members) string {
	var query string = "id = "
	if len(members) == 1{
		query = fmt.Sprintf("%s '%s'", query, members[0].OrganisationID)
		return query
	}
	for i := 0; i < len(members); i++ {
		if i == len(members)-1 {
			query = fmt.Sprintf("%s '%s'", query, members[i].OrganisationID)
		} else {
			query = fmt.Sprintf("%s '%s' OR id =", query, members[i].OrganisationID)
		}
	}
	return query
}
