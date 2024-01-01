package permission_groups

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/permission_group"
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
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.PermissionGroupModule()))
	response["data"] = transformers.PermissionGroupsResponse(rows, false)
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
	helpers.OkResponse(helpers.DoneCreateItem(models.PermissionGroupModel()), storeUpdateData(&row))
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
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), transformers.PermissionGroupResponse(row))
}

func storeUpdateData(row interface{}) map[string]interface{} {
	/// convert interface to row
	category := row.(*models.PermissionGroup)
	/// transform row
	m := transformers.PermissionGroupResponse(*category)
	/// store translate rows
	translateRows := models.StoreTranslateColumn(m, models.PermissionGroupTransTable(), models.PermissionGroupTranslateFillable())
	/// transform translate rows
	m = transformers.TransformTranslation(translateRows, m)

	return m
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
	_const.Services.DB.Unscoped().Delete(&row)
	helpers.DeleteRowsTranslations(models.CategoryTransTable(), row.ID)
	// now return ok response
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.PermissionGroupModel()))
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
	inter := models.Store(models.PermissionGroupModule(), models.PermissionGroupFillAbleColumn(), &oldRow)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.PermissionGroupModel()), storeUpdateData(inter))
}

/***
* quick edit
 */
func QuickEdit(g *gin.Context) {
	///validate request
	err := permission_group.QuickEdit()
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
	helpers.OkResponse(helpers.DoneUpdate(models.PermissionGroupModel()), storeUpdateData(&oldRow))
}
