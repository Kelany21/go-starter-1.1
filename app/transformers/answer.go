package transformers

import (
	"starter-golang-new/app/models"
	"starter-golang-new/helpers"
)

/**
* stander the single Answer response
 */
func AnswerResponse(answer models.Answer) map[string]interface{} {
	var u = make(map[string]interface{})
	u["id"] = answer.ID
	u["faq_id"] = answer.FaqId

	u["created_at"] = helpers.StringDateReformat(answer.CreatedAt)
	u["updated_at"] = helpers.StringDateReformat(answer.UpdatedAt)
	u["created_time"] = helpers.StringTimeReformat(answer.CreatedAt)
	u["updated_time"] = helpers.StringTimeReformat(answer.UpdatedAt)

	return u
}

/**
* stander the Multi Answers response
 */
func AnswersResponse(answers []models.Answer, withLang bool) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	translations := GetTranslations(helpers.GetIDs(answers), models.AnswerTransTable())
	for _, answer := range answers {
		if withLang {
			u = append(u, TransformTranslationObject(translations[uint64(answer.ID)], AnswerResponse(answer)))
		} else {
			u = append(u, TransformTranslation(translations[uint64(answer.ID)], AnswerResponse(answer)))
		}
	}
	return u
}
