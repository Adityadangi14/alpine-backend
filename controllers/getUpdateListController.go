package controllers

import (
	"net/http"
	initializers "project_mine/initlizers"
	"project_mine/model"

	"github.com/gin-gonic/gin"
)

func GetUpdateListController(c *gin.Context) {
	var tableList []model.Table

	row := initializers.DB.Order("created_at asc").Find(&tableList).RowsAffected

	if row == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Unable to fetch data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tableList,
	})

}
