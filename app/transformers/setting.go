package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single setting response
 */
func SettingResponse(setting models.Setting) map[string]interface{} {
	var u = make(map[string]interface{})
	u["value"] = setting.Value
	u["id"] = setting.ID
	u["name"] = setting.Name
	u["slug"] = setting.Slug
	u["setting_type"] = setting.SettingType
	u["created_at"] = helpers.StringDateReformat(setting.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(setting.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(setting.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(setting.UpdatedAt)


	return u
}

/**
* stander the Multi settings response
 */
func SettingsResponse(settings []models.Setting) []map[string]interface{} {
	var u  = make([]map[string]interface{} , 0)
	for _ , setting := range settings {
		u = append(u , SettingResponse(setting))
	}
	return u
}

