package initializers

import "project_mine/model"

func SyncDB() {
	DB.Exec("DEALLOCATE PREPARE ALL;")
	DB.AutoMigrate(&model.Table{})
	DB.AutoMigrate(&model.UserModel{})
	DB.AutoMigrate(&model.NotficationPool{})
}
