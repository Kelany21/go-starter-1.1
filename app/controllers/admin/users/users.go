package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/user"
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
	response := make(map[string]interface{})
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.UserModule()))
	response["data"] = transformers.UsersResponse(rows)
	// transform slice
	paginator.Records = response
	// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
}

/***
get all status
*/
func Status(g *gin.Context) {
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.ActionsResponse(models.GetActionByModule(models.UserModule())))
}

/**
* store new user
 */
func Store(g *gin.Context) {
	// check if request valid
	if !validateRequest("store") {
		return
	}
	row := ConvertRequestRow()
	/// check if this email exists
	// create new row
	_const.Services.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneCreateItem(models.UserModel()), transformers.UserResponse(row))
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
	helpers.OkResponse(helpers.DoneGetItem(), transformers.UserResponse(row))
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	authUser := models.Auth()
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	if authUser.ID == row.ID {
		helpers.YouCantTrashOrDeleteYourSelf()
		return
	}
	_const.Services.DB.Unscoped().Delete(&row)
	// now return row data after transformers
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.UserModel()))
}

/**
* update user
 */
func Update(g *gin.Context) {
	// check if request valid
	if !validateRequest("update") {
		return
	}
	// find this row or return 404
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	if _, ok := d["password"]; ok {
		if d["password"].(string) != ""{
			d["password"], _ = helpers.HashPassword(d["password"].(string))
		}
	}
	inter := models.UpdateTest(models.UserTable(), models.UserFillAbleColumn(), &oldRow, oldRow.ID, d)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.UserModel()), transformers.UserResponse(*inter.(*models.User)))
}

/***
* quick edit
 */
func QuickEdit(g *gin.Context) {
	///validate request
	err := user.QuickEdit()
	/// return if error
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	oldRow := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &oldRow); oldRow.ID == 0 {
		return
	}
	inter := models.UpdateTest(models.UserTable(), models.UserFillAbleColumn(), &oldRow, oldRow.ID)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.UserModel()), transformers.UserResponse(*inter.(*models.User)))
}
