package seeders

import (
	"os"
	"strconv"
)

func Seed() {
	ifDrop, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
	if ifDrop {
		StatusSeeder()
		CategorySeeder()
		PermissionGroupSeeder()
		UserSeeder()
		SettingSeeder()
		PageSeeder()
		RoleSeeder()
	}
}
