package utils

import (
	"fmt"
	"strings"
)

func OrgName(userName string) string {
	var orgName string
	if _, found := strings.CutSuffix(strings.ToLower(userName), "s"); found {
		orgName = fmt.Sprintf("%s' Organisation", userName)
	} else {
		orgName = fmt.Sprintf("%s's Organisation", userName)
	}
	return orgName
}
