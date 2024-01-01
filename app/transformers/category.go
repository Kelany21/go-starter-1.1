package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single user response
 */
func CategoryResponse(category models.Category) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = category.ID
	u["status"] = category.Status

	u["created_at"] = helpers.StringDateReformat(category.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(category.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(category.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(category.UpdatedAt)

	return u
}

/**
* stander the Multi categories response
 */
func CategoriesResponse(categories []models.Category, withLang bool) []map[string]interface{} {
	translations := GetTranslations(helpers.GetIDs(categories), models.CategoryTransTable())
	var u = make([]map[string]interface{}, 0)
	for _, category := range categories {
		if withLang {
			u = append(u, TransformTranslationObject(translations[uint64(category.ID)], CategoryResponse(category)))
		} else {
			u = append(u, TransformTranslation(translations[uint64(category.ID)], CategoryResponse(category)))
		}
	}
	return u
}
