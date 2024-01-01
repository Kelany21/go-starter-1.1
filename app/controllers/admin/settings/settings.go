package settings

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"starter-golang-new/app/models"
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
	response["status"] = transformers.ActionsResponse(models.GetActionByModule(models.SettingModule()))
	response["data"] = transformers.SettingsResponse(rows)
	// transform slice
	paginator.Records = response
	// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
}

/***
get all status
*/
func Status(g *gin.Context) {
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.ActionsResponse(models.GetActionByModule(models.SettingModule())))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	// find this row or return 404
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), transformers.SettingResponse(row))
}

/**
* update category
 */
func Update(g *gin.Context) {
	// check if request valid
	if !validateRequest() {
		return
	}
	_const.Services.DB.Unscoped().Delete(&models.Setting{})
	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	for name, value := range d {
		_const.Services.DB.Model(&models.Setting{}).Where("slug = ?", name).Updates(map[string]interface{}{"value": value})
	}
	allRows := Rows()
	_const.Services.DB.Find(&allRows)
	//now return row data after transformers
	helpers.OkResponse(helpers.DoneUpdate(models.CategoryModel()), transformers.SettingsResponse(allRows))
}

func MakeBackup(c *gin.Context) {
	cmd, er := exec.Command("mysqldump", "-u", os.Getenv("DATABASE_USERNAME"), "-p" + os.Getenv("DATABASE_PASSWORD"), "--databases", os.Getenv("DATABASE_NAME")).Output()

	if  er != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "test",
				"data":    er.Error(),
				//"data":    string(er.(*exec.ExitError).Stderr),
			})
		return
	}
	f, err := os.Create("test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(string(cmd))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	c.JSON(
		http.StatusBadRequest, gin.H{
			"status":  true,
			"message": "done",
			"data":    "",
		})
	return
}
