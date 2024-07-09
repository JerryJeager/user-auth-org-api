package utils

import "github.com/JerryJeager/user-auth-org-api/internal/service/models"

//returns true if two members are in the same organisation

func IsInSameOrganisation (user1, user2 models.Members) bool {
	for i := 0; i < len(user1); i++{
		for j := 0; j < len(user2); j++ {
			if user1[i].OrganisationID.String() == user2[j].OrganisationID.String() {
				return true
			}
		}
	}
	return false
}