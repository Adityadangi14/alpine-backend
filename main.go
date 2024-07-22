package main

import (
	"project_mine/controllers"
	initializers "project_mine/initlizers"
	loghandler "project_mine/logHandler"
	"project_mine/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	loghandler.OpenLogFile()
	loghandler.Log()
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDB()
	controllers.PushUpdates()
}
func main() {
	r := gin.Default()

	r.GET("/getTableList", middleware.RequiredAuth, controllers.GetUpdateListController)
	r.POST("/getStockUrl", middleware.RequiredAuth, controllers.GetStockUrl)
	r.POST("/auth", controllers.AuthHandler)

	err := r.Run(":3000")

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
	}
}
