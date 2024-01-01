package faqs

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* loop on the answers
* insert all answer in database
 */
func insertAnswerToDataBase(id uint) {
	request := models.ConvertBodyToHashMap()
	answers := request["answers"].([]interface{})
	if len(answers) > 0 {
		for i, _ := range answers {
			answer := answers[i].(map[string]interface{})
			texts := answer["text"].([]interface{})
			row := models.Answer{FaqId: int(id)}
			_const.Services.DB.Create(&row)
			_const.Services.DB.Create(&models.AnswerI18ns{ReferenceId: uint64(row.ID), FiledName: "text", Language: "ar", Value: texts[0].(string),})
			_const.Services.DB.Create(&models.AnswerI18ns{ReferenceId: uint64(row.ID), FiledName: "text", Language: "en", Value: texts[1].(string),})
		}
	}
	return
}

/**
*  delete answers
 */
func deleteAnswers(faqId uint) {
	var ids []int
	_const.Services.DB.Model(models.Answer{}).Where("faq_id = ? ", faqId).Pluck("id", &ids)
	helpers.DeleteRowsTranslations(models.AnswerTransTable(), helpers.ConvertArrayIntToArrayInterfaces(ids)...)
	var rows models.Answer
	_const.Services.DB.Unscoped().Where("faq_id = ? ", faqId).Delete(&rows)
}

/**
* store new faq
 */
func StoreAnswer(g *gin.Context) {
	/// check if request valid
	if !validateRequestAnswer() {
		return
	}
	/// create new row
	row := ConvertRequestRowAnswer()
	_const.Services.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.AnswerModel()), storeUpdateDataAnswer(&row))
}

func UpdateAnswer(g *gin.Context) {
	/// check if request valid
	if !validateRequestAnswer() {
		return
	}
	oldRow := RowAnswer()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	inter := models.UpdateTest(models.AnswerTable(), models.AnswerFillAbleColumn(), &oldRow, oldRow.ID)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.AnswerModel()), storeUpdateDataAnswer(inter))
}

func storeUpdateDataAnswer(row interface{}) map[string]interface{} {
	/// convert interface to row
	answer := row.(*models.Answer)
	/// transform row
	m := transformers.AnswerResponse(*answer)
	/// store translate rows
	translateRows := models.StoreTranslateColumn(m, models.AnswerTransTable(), models.AnswerTranslateFillable())
	/// transform translate rows
	m = transformers.TransformTranslation(translateRows, m)

	return m
}
