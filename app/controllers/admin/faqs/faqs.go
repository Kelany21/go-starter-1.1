package faqs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/faq"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* get all rows with pagination
 */
func Index(g *gin.Context) {
	/// array of rows
	rows := Rows()
	/// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      _const.Services.DB,
		Page:    helpers.Page(),
		Limit:   helpers.Limit(),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	/// transform slice
	response := make(map[string]interface{})
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.FaqModule()))
	response["data"] = transformers.FaqsResponse(rows, false)
	// transform slice
	paginator.Records = response
	/// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
}

/**
* store new faq
 */
func Store(g *gin.Context) {
	/// check if request valid
	if !validateRequest() {
		return
	}
	rowMap, _ := store()
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.FaqModel()), rowMap)
}

/**
* store new faq with answers
 */
func StoreFaqWithAnswers(g *gin.Context) {
	err := faq.StoreUpdateFaqWithArray()
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	_, row := store()
	// insert answers
	insertAnswerToDataBase(row.ID)
	models.FindOrFail(row.ID, &row, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	})
	/// transform row
	m := transformers.FaqResponse(row, false)
	/// transform translation
	m = transformers.TransformTranslation(helpers.RowsTranslations(models.FaqTransTable(), row.ID), m)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.FaqModel()), m)
} /**
* store new faq with answers
 */
func UpdateFaqWithAnswers(g *gin.Context) {
	err := faq.StoreUpdateFaqWithArray()
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	// update faq
	inter := models.UpdateTest(models.FaqTable(), models.FaqFillAbleColumn(), &oldRow, oldRow.ID)
	// update translations
	storeUpdateData(inter)
	// delete old answers
	deleteAnswers(oldRow.ID)
	// insert new answers
	insertAnswerToDataBase(oldRow.ID)
	// find after update
	models.FindOrFail(oldRow.ID, &oldRow, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	})
	/// transform row
	m := transformers.FaqResponse(oldRow, false)
	/// transform translation
	m = transformers.TransformTranslation(helpers.RowsTranslations(models.FaqTransTable(), oldRow.ID), m)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.FaqModel()), m)
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	}); row.ID == 0 {
		return
	}
	/// transform row
	m := transformers.FaqResponse(row, false)
	/// transform translation
	m = transformers.TransformTranslation(helpers.RowsTranslations(models.FaqTransTable(), row.ID), m)
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), m)

}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	/// find this row or return 404
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	/// delete related answers
	deleteAnswers(row.ID)
	/// delete row
	_const.Services.DB.Unscoped().Delete(&row)
	// delete translations
	helpers.DeleteRowsTranslations(models.FaqTransTable(), row.ID)
	/// now return ok response
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.FaqModel()))
}

/**
* update faq
 */
func Update(g *gin.Context) {
	/// check if request valid
	if !validateRequest() {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	inter := models.UpdateTest(models.FaqTable(), models.FaqFillAbleColumn(), &oldRow, oldRow.ID)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.FaqModel()), storeUpdateData(inter))
}

/***
* quick edit
 */
func QuickEdit(g *gin.Context) {
	err := faq.QuickEdit()
	/// return if error
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	/// update allow columns and return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.FaqModel()), storeUpdateData(&oldRow))
}

func storeUpdateData(row interface{}) map[string]interface{} {
	/// convert interface to row
	faq := row.(*models.Faq)
	/// transform row
	m := transformers.FaqResponse(*faq, false)
	/// store translate rows
	translateRows := models.StoreTranslateColumn(m, models.FaqTransTable(), models.FaqTranslateFillable())
	/// transform translate rows
	m = transformers.TransformTranslation(translateRows, m)

	return m
}
