package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	initializers "project_mine/initlizers"
	loghandler "project_mine/logHandler"
	"project_mine/model"
	"project_mine/service"
	"time"
)

func PushUpdates() {
	var tableData model.Table

	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {

			data, err := getDataFromApi()
			var wg sync.WaitGroup

			if err != nil {
				loghandler.AppLogger.Error(string(err.Error()))
			}

			for _, table := range data {
				wg.Add(1)
				initializers.DB.Where("news_id = ?", table.NEWSID).First(&tableData)
				wg.Done()
				fmt.Println(tableData.NEWSID)
				if tableData.NEWSID == "" {
					initializers.DB.Create(&table)
					service.PushNotificationService(table.SLONGNAME+"/ "+table.CATEGORYNAME, table.HEADLINE)
				}

			}

		}
	}()

}

func getDataFromApi() ([]model.Table, error) {
	baseUrl := "https://api.bseindia.com/BseIndiaAPI/api/AnnSubCategoryGetData/w"
	params := url.Values{
		"pageno":      {"1"},
		"strCat":      {"Company Update"},
		"strPrevDate": {getCurrentDateYYYYMMDD()},
		"strScrip":    {""},
		"strSearch":   {"P"},
		"strToDate":   {getCurrentDateYYYYMMDD()},
		"strType":     {"C"},
		"subcategory": {"Award of Order / Receipt of Order"},
	}
	url := baseUrl + "?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)

	fmt.Println(url)

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		return []model.Table{}, err
	}

	req.Header.Add("Referer", "https://www.bseindia.com/")
	req.Header.Add("Origin", "https://www.bseindia.com/")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.3")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")

	fmt.Println(req.Header.Get("Referer"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		return []model.Table{}, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		return []model.Table{}, err
	}

	type Data struct {
		Table  []model.Table `json:"Table"`
		Table1 []struct {
			ROWCNT int `json:"ROWCNT"`
		} `json:"Table1"`
	}
	var data Data

	er := json.Unmarshal(content, &data)

	if er != nil {
		return []model.Table{}, er
	}

	return data.Table, nil
}
func getCurrentDateYYYYMMDD() (date string) {
	now := time.Now()

	// Format the date in the desired format (YYYYMMDD)
	formattedDate := now.Format("20060102")

	return formattedDate
}
