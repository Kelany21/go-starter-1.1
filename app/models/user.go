package models

import (
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type User struct {
	gorm.Model
	Name              string `gorm:"type:varchar(50);" json:"name"`
	Email             string `gorm:"type:varchar(50);unique_index" json:"email"`
	PermissionGroupId int    `gorm:"_" json:"permission_group_id"`
	Password          string `gorm:"size:255" json:"password"`
	Token             string `gorm:"size:255" json:"token"`
	Status            string `gorm:"type:varchar(20);" json:"status"`
	Image             string `gorm:"size:255" json:"image"`
}

/**
* use this struct when visitor login
 */
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

/**
* use this struct when reset email
 */
type Reset struct {
	Email string `json:"email"`
}

/**
* use this struct when reset email
 */
type Recover struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

/*
* return with module name
 */
func UserModule() string {
	return helpers.ModuleName("users")
}

func UserTable() string {
	return helpers.ModuleTable(UserModule())
}

func UserTransTable() string {
	return helpers.ModuleTransTable(UserModule())
}

/*
* return with model name
 */
func UserModel() string {
	return helpers.ModelModel(UserModule())
}

/*
* return with route name
 */
func UserRoute() string {
	return helpers.ModuleRoute(UserModule())
}

/*
* return with Statuses
 */
func UserStatuses() interface{} {
	return helpers.ModuleStatuses(UserModule())
}

/**
* event when user register
* create token
* hash password
 */
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	if user.Email != "admin@admin.com" {
		token, _ := helpers.HashPassword(user.Email + user.Password)
		scope.SetColumn("token", token)
	}
	password, _ := helpers.HashPassword(user.Password)
	scope.SetColumn("password", password)

	return nil
}

/*
* event run after user register
 */
func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(UserModule())
	return
}

/*
* event run after delete user
 */
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, UserModule())
	return
}

/**
* update status
 */
func (u *User) BeforeUpdate() (err error) {
	var user User
	if u.ID != 0 {
		_const.Services.DB.First(&user, u.ID)
		if user.Status != u.Status {
			if u.Status == _const.TRASH {
				DecreaseRow(_const.ALL, UserModule())
			}
			if user.Status == _const.TRASH {
				IncreaseRow(_const.ALL, UserModule())
			}
			DecreaseRow(user.Status, UserModule())
			IncreaseRow(u.Status, UserModule())
		}
	}
	return
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func UserMigrate() {
	_const.Services.DB.AutoMigrate(&User{})
}

/**
* you can update these column only
 */
func UserFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"
	m["email"] = "string"
	m["permission_group_id"] = "number"
	m["password"] = "string"
	m["status"] = "string"
	m["image"] = "string"

	return m
}

func UserTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"
	m["description"] = "text"

	return m
}

/**
* active category only
 */
func ActiveUser(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + _const.ACTIVE)
}

/**
* get auth user
 */
func Auth() User {
	var user User
	_const.Services.DB.Where("token = ? and status = ?", helpers.GetClearToken(), _const.ACTIVE).First(&user)
	return user
}

func GetUserTable() Table {
	table := Table{
		Name: "users",
		Columns: []Column{DefulteColumn("name"), {
			Name:       "block",
			Sort:       true,
			Show:       true,
			Label:      "User Status",
			RenderType: "text",
			Align:      "center",
			Filter: Filter{
				ShowFilter:         true,
				FilterType:         "select",
				DefaultFilterValue: "",
				FilterOptions: []Option{{
					Text: "select user status",
				}, {
					Text:  "Block",
					Value: 1,
				}, {
					Text:  "Active",
					Value: 2,
				}},
			},
			Form: Form{
				InputType:      "select",
				SubmitOnUpdate: true,
				SubmitOnCreate: true,
				StoreType:      "integer",
				Placeholder:    "select status",
				QuickEdit:      true,
			},
		}},
	}
	return table
}
