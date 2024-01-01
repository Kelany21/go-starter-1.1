package seeders

import (
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"syreclabs.com/go/faker"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func UserSeeder() {
	data := models.User{
		Email:             "admin@admin.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "abeer@abeer.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "bassem@bassem.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "radwa@radwa.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "nasser@nasser.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "hams@hams.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	data = models.User{
		Email:             "adham@adham.com",
		Password:          "123456",
		Token:             "$2a$10$PVD8Iu0sLrojlg7mzjNsLu04zs.hFBnDpLPOGgJeX2nb7IE62kS42",
		Name:              "Admin",
		Status:            _const.ACTIVE,
		PermissionGroupId: _const.ADMIN_ID,
	}
	_const.Services.DB.Save(&data)
	newUser(true)
	for i := 0; i < 10; i++ {
		newUser(false)
	}
}

/**
* fake data and create data base
 */
func newUser(admin bool) {
	data := models.User{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(8, 14),
		Name:     faker.Internet().UserName(),
		Status:   _const.ACTIVE,
	}
	if admin {
		data.PermissionGroupId = _const.ADMIN_ID
	} else {
		data.PermissionGroupId = _const.USER_ID
	}
	_const.Services.DB.Create(&data)
}
