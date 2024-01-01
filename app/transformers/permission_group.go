package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single image page response
 */
func PermissionGroupResponse(permissionGroup models.PermissionGroup) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = permissionGroup.ID
	u["status"] = permissionGroup.Status

	u["created_at"] = helpers.StringDateReformat(permissionGroup.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(permissionGroup.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(permissionGroup.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(permissionGroup.UpdatedAt)
	return u
}

/**
* stander the Multi images page response
 */
func PermissionGroupsResponse(permissionGroups []models.PermissionGroup, withLang bool) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	translations := GetTranslations(helpers.GetIDs(permissionGroups), models.CategoryTransTable())
	for _, permissionGroup := range permissionGroups {
		if withLang {
			u = append(u, TransformTranslationObject(translations[uint64(permissionGroup.ID)], PermissionGroupResponse(permissionGroup)))
		} else {
			u = append(u, TransformTranslation(translations[uint64(permissionGroup.ID)], PermissionGroupResponse(permissionGroup)))
		}
	}
	return u
}
