package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single faq response
 */
func FaqResponse(faq models.Faq, withLang bool) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = faq.ID
	u["status"] = faq.Status
	u["answer"] = AnswersResponse(faq.Answers, withLang)

	u["created_at"] = helpers.StringDateReformat(faq.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(faq.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(faq.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(faq.UpdatedAt)

	return u
}

/**
* stander the Multi faqs response
 */
func FaqsResponse(faqs []models.Faq, withLang bool) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	translations := GetTranslations(helpers.GetIDs(faqs), models.FaqTransTable())
	for _, faq := range faqs {
		if withLang {
			u = append(u, TransformTranslationObject(translations[uint64(faq.ID)], FaqResponse(faq, true)))
		} else {
			u = append(u, TransformTranslation(translations[uint64(faq.ID)], FaqResponse(faq, false)))
		}
	}
	return u
}
