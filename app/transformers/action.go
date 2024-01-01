package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single user response
 */
func ActionResponse(status models.Status) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = status.ID
	u["noun"] = status.Noun
	u["verb"] = status.Verb
	u["slug"] = status.Slug
	u["module_name"] = status.ModuleName
	u["slug"] = status.Slug
	u["count"] = status.Count
	u["created_at"] = helpers.StringDateReformat(status.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(status.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(status.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(status.UpdatedAt)

	return u
}

/**
* stander the Multi users response
 */
func ActionsResponse(statuses []models.Status) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, status := range statuses {
		u = append(u, ActionResponse(status))
	}
	return u
}
