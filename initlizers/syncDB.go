package initializers

import (
	"project_mine/model"
)

func SyncDB() {
	DB.AutoMigrate(&model.Table{})
	DB.AutoMigrate(&model.UserModel{})

}
