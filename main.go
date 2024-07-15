package main

import (
	"fmt"
	"project_mine/controllers"
	initializers "project_mine/initlizers"
	"project_mine/middleware"

	"github.com/gin-gonic/gin"
)

func init() {

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
		fmt.Println(err)
	}
}
