package categories

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/category"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)


/***
* get all rows with pagination
 */
func Index(g *gin.Context) {
	// array of rows
	rows := Rows()
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      _const.Services.DB,
		Page:    helpers.Page(),
		Limit:   helpers.Limit(),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)

	// transform slice
	response := make(map[string]interface{})
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.CategoryRoute()))
	response["data"] = transformers.CategoriesResponse(rows, false)
	// transform slice
	paginator.Records = response
	// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)

}


/**
* store new category
 */
func Store(g *gin.Context) {
	// check if request valid
	if !validateRequest() {
		return
	}
	/// store row and transform return row
	row := ConvertRequestRow()
	_const.Services.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.CategoryModel()), storeUpdateData(&row))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	/// transform row
	m := transformers.CategoryResponse(row)
	/// transform translation
	m = transformers.TransformTranslation(helpers.RowsTranslations(models.CategoryTransTable(), row.ID), m)
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), m)
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	/// delete main row
	_const.Services.DB.Unscoped().Delete(&row)
	/// delete translate rows
	helpers.DeleteRowsTranslations(models.CategoryTransTable(), row.ID)
	// now return ok response
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.CategoryModel()))
}

/**
* update category
 */
func Update(g *gin.Context) {
	// check if request valid
	if !validateRequest() {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	inter := models.UpdateTest(models.CategoryTable(), models.CategoryFillAbleColumn(), &oldRow, oldRow.ID)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.CategoryModel()), storeUpdateData(inter))
}

func storeUpdateData(row interface{}) map[string]interface{} {
	/// convert interface to row
	category := row.(*models.Category)
	/// transform row
	m := transformers.CategoryResponse(*category)
	/// store translate rows
	translateRows := models.StoreTranslateColumn(m, models.CategoryTransTable(), models.CategoryTranslateFillable())
	/// transform translate rows
	m = transformers.TransformTranslation(translateRows, m)

	return m
}

/***
* quick edit
 */
func QuickEdit(g *gin.Context) {
	///validate request
	err := category.QuickEdit()
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
	helpers.OkResponse(helpers.DoneUpdate(models.CategoryModel()), storeUpdateData(&oldRow))
}
