package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single page response
 */
func PageResponse(page models.Page) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = page.ID
	u["status"] = page.Status
	u["slug"] = page.Slug
	u["images"] = PageImagesResponse(page.Images)

	u["created_at"] = helpers.StringDateReformat(page.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(page.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(page.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(page.UpdatedAt)

	return u
}

/**
* stander the Multi pages response
 */
func PagesResponse(pages []models.Page, withLang bool) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	translations := GetTranslations(helpers.GetIDs(pages), models.PageTransTable())
	for _, page := range pages {
		if withLang {
			u = append(u, TransformTranslationObject(translations[uint64(page.ID)], PageResponse(page)))
		} else {
			u = append(u, TransformTranslation(translations[uint64(page.ID)], PageResponse(page)))
		}
	}
	return u
}
