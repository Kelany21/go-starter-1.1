package pages

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/page"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* get all rows with pagination
 */
func Index(g *gin.Context) {
	// array of rows
	var rows []models.Page
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
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.PageModule()))
	response["data"] = transformers.PagesResponse(rows, false)
	// transform slice
	paginator.Records = response
	// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
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
	/// get all page images
	getImages(&row)
	/// transform row
	m := transformers.PageResponse(row)
	/// transform translation
	m = transformers.TransformTranslation(helpers.RowsTranslations(models.PageTransTable(), row.ID), m)
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), m)
	// now return row data after transformers
}

/**
* update page
 */
func Store(g *gin.Context) {
	// check if request valid
	if ! validateRequest() {
		return
	}
	// find this row or return 404
	row := ConvertRequestRow()
	
	inter := models.Store(models.PageTable(), models.PageFillAbleColumn(), &row)
	/// upload images
	insertImageInDataBase(getImageFromRequest(), int(row.ID))

	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.PageModel()), storeUpdateData(inter))

}

/**
* update page
 */
func Update(g *gin.Context) {
	// check if request valid
	if ! validateRequest() {
		return
	}
	// find this row or return 404
	var oldRow models.Page
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	/// delete all images if reset flag in the url
	if g.Query("reset") == "true" {
		deleteAllPageImage(oldRow.ID)
	}

	inter := models.UpdateTest(models.PageTable(), models.PageFillAbleColumn(), &oldRow, oldRow.ID)
	/// upload images
	insertImageInDataBase(getImageFromRequest(), int(oldRow.ID))

	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.PageModel()), storeUpdateData(inter))

}

/***
* quick edit
 */
func QuickEdit(g *gin.Context) {
	///validate request
	err := page.QuickEdit()
	/// return if error
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	// update allow columns and now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.PageModel()), storeUpdateData(&oldRow))
}

func storeUpdateData(row interface{}) map[string]interface{} {
	/// convert interface to row
	p := row.(*models.Page)
	/// transform row
	getImages(p)
	m := transformers.PageResponse(*p)
	/// store translate rows
	translateRows := models.StoreTranslateColumn(m, models.PageTransTable(), models.PageTranslateFillable())
	/// transform translate rows
	m = transformers.TransformTranslation(translateRows, m)

	return m
}

func getImageFromRequest() []string {
	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	var arr []interface{}
	var images []string
	arr = d["image"].([]interface{})
	for _, image := range arr {
		images = append(images, image.(string))
	}
	return images
}
