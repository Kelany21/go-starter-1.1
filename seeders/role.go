package seeders

import (
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func RoleSeeder() {
	var permissionGroups []models.PermissionGroup
	_const.Services.DB.Find(&permissionGroups)
	for _, permissionGroup := range permissionGroups {
		for moduleName, module := range _const.Services.Modules {
			modulePermission := module["permissions"].(map[string]interface{})
			for slug, functions := range modulePermission {
				for _, function := range functions.([]interface{}) {
					role := models.Role{
						Slug:              slug,
						FuncName:          function.(string),
						PermissionGroupId: int(permissionGroup.ID),
						Module:            moduleName,
					}
					_const.Services.DB.Create(&role)
				}
			}
		}
	}
}
