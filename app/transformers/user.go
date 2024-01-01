package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single user response
 */
func UserResponse(user models.User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = user.ID
	u["name"] = user.Name
	u["email"] = user.Email
	u["image"] = user.Image
	u["permission_group_id"] = user.PermissionGroupId
	u["token"] = user.Token
	u["status"] = user.Status
	u["created_at"] = helpers.StringDateReformat(user.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(user.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(user.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(user.UpdatedAt)

	return u
}

/**
* stander the Multi users response
 */
func UsersResponse(users []models.User) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, user := range users {
		u = append(u, UserResponse(user))
	}
	return u

}
