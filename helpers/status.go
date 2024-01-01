package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/url"
	"reflect"
	_const "starter-golang-new/const"
	"strconv"
)

/**
* status that make item active
 */
func Active(id interface{}, object interface{}) {
	changeStatus(id, _const.ACTIVE, object)
}

/**
* status that make item deactive
 */
func DeActive(id interface{}, object interface{}) {
	changeStatus(id, _const.DEACTIVATE, object)
}

/**
* status that make item trash
 */
func Trash(id interface{}, object interface{}) {
	changeStatus(id, _const.TRASH, object)
}

/**
* status that make item block
 */
func Block(id interface{}, object interface{}) {
	changeStatus(id, _const.BLOCK, object)
}

/**
* status that make item active
 */
func BulkActive(ids interface{}, model interface{}) {
	bulkChangeStatus(ids, _const.ACTIVE, model)
}

/**
* status that make item deactive
 */
func BulkDeActive(ids interface{}, model interface{}) {
	bulkChangeStatus(ids, _const.DEACTIVATE, model)
}

/**
* status that make item trash
 */
func BulkTrash(ids interface{}, model interface{}) {
	bulkChangeStatus(ids, _const.TRASH, model)
}

/**
* status that make item block
 */
func BulkBlock(ids interface{}, model interface{}) {
	bulkChangeStatus(ids, _const.BLOCK, model)
}

/**
* delete items
 */
func BulkDelete(ids interface{}, model interface{}) {
	bulkChangeStatus(ids, _const.DELETE, model)
}

/**
* delete item
 */
func GetStatusSeparateWithComma() string {
	return _const.ACTIVE + "," + _const.DEACTIVATE + "," + _const.TRASH
}

func BulkInit(action string, object interface{}) {
	var ids _const.IDS
	err := ValidationBulk()
	if ReturnNotValidRequest(err) {
		return
	}
	ids = ReadRequest(object)
	switch action {
	case _const.ACTIVE:
		BulkActive(ids.Ids, object)
		break
	case _const.DEACTIVATE:
		BulkDeActive(ids.Ids, object)
		break
	case _const.TRASH:
		BulkTrash(ids.Ids, object)
		break
	case _const.BLOCK:
		BulkBlock(ids.Ids, object)
		break
	case _const.DELETE:
		BulkDelete(ids.Ids, object)
		break
	}
	return
}

func ReadRequest(object interface{}) _const.IDS {
	ids := _const.IDS{}
	if err := json.Unmarshal(_const.GetBodyB(), &ids); err != nil {
		fmt.Println(err)
	}

	if GetTableNameFormStruct(object) == "users" {
		userId, _ := strconv.Atoi(_const.Services.GIN.Request.Header.Get("user_id"))
		var ides []int
		for i := 0; i < len(ids.Ids); i++ {
			if ids.Ids[i] != 1 && ids.Ids[i] != userId {
				ides = append(ides, ids.Ids[i])
			}
		}
		ids.Ids = ides
	}
	return ids
}

/***
* change row status based on status
* return 404 if not found
* return done if every thing is valid
 */
func changeStatus(id interface{}, status string, object interface{}) {
	/// check if row found
	_const.Services.DB.Where("id = ? ", id).First(object)
	findId := reflect.ValueOf(object).Elem().FieldByName("ID").Uint()
	if findId == 0 {
		ReturnNotFound(ItemNotFound())
		_const.Services.GIN.Abort()
		return
	}
	/// update status
	_const.Services.DB.Model(object).Update("status", status)
	switch status {
	case _const.ACTIVE:
		ActionItem(T("item_has_been_activated"))
		break
	case _const.DEACTIVATE:
		ActionItem(T("item_has_been_deactivated"))
		break
	case _const.TRASH:
		ActionItem(T("item_has_been_trashed"))
		break
	case _const.BLOCK:
		ActionItem(T("item_has_been_block"))
		break
	}

	return
}

type De struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

/***
* update new ids with new status
 */
func bulkChangeStatus(ids interface{}, status string, model interface{}) {

	length := len(ids.([]int))
	moduleName := GetTableNameFormStruct(model)
	var des []De
	/// group ids with status first
	/// then update all status by looping
	_const.Services.DB.Table(moduleName).Select("status , COUNT(id) as count").Where("id IN(?)", ids).Group("status").Find(&des)
	for _, de := range des {
		Decrease("statuses", "count", nil, strconv.Itoa(de.Count), ` module_name = "`+moduleName+`"`, ` verb = "`+de.Status+`"`)
	}
	/// if action delete Decrease from all
	if _const.DELETE != status {
		/// update status
		_const.Services.DB.Model(model).Where("id IN(?)", ids).Update("status", status)
		/// then increase with all number to the new one
		Increase("statuses", "count", nil, strconv.Itoa(length), ` module_name = "`+moduleName+`"`, ` verb = "`+status+`"`)
	} else {
		Decrease("statuses", "count", nil, strconv.Itoa(length), ` module_name = "`+moduleName+`"`, ` verb = "all"`)
	}
	switch status {
	case _const.ACTIVE:
		ActionItem(T("item_has_been_activated"))
		break
	case _const.DEACTIVATE:
		ActionItem(T("item_has_been_deactivated"))
		break
	case _const.TRASH:
		ActionItem(T("item_has_been_trashed"))
		break
	case _const.BLOCK:
		ActionItem(T("item_has_been_block"))
		break
	case _const.DELETE:
		ActionItem(T("item_has_been_deleted"))
		break
	}

	return
}

/**
* exception validation for bulk action only
* any other validation will store in requests
* folder
 */
func ValidationBulk() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"ids": []string{"int_slice"},
	}
	messages := govalidator.MapData{
		"ids": []string{IntSlice()},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request:  _const.Request(), // request object
		Rules:    rules,            // rules map
		Messages: messages,         // custom message map (Optional)
		Data:     &data,
	}

	vd := govalidator.New(opts)

	return vd.ValidateJSON()
}
