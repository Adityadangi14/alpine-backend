package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	loghandler "project_mine/logHandler"

	"github.com/gin-gonic/gin"
)

func GetStockUrl(c *gin.Context) {
	var body struct {
		StockName string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "unbable to read body",
		})
	}

	type UrlData []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	res, er := http.Get("https://www.screener.in/api/company/search/?q=" + body.StockName + "&v=3&fts=1")

	if er != nil {

		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "unbable to get Url ",
		})
	}

	defer res.Body.Close()

	content, e := io.ReadAll(res.Body)

	if e != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "unbable to get Url ",
		})
	}

	var data UrlData

	err := json.Unmarshal(content, &data)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "unbable to get Url ",
		})

		loghandler.AppLogger.Error(string(err.Error()))

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "https://www.screener.in" + data[0].URL,
	})

}
