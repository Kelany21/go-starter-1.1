package transformers

import (
	"starter-golang-new/app/models"
	"strings"
)

/**
* stander the Multi images page response
 */
func RolesResponse(roles []models.Role) []string {
	var u []string
	for _, role := range roles {
		u = append(u, strings.ToLower(role.Module + "_" + role.FuncName))
	}
	return u
}
