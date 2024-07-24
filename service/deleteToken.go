package service

import (
	"fmt"
	initializers "project_mine/initlizers"
	loghandler "project_mine/logHandler"
	"project_mine/model"
	"slices"

	"github.com/lib/pq"
)

func DeleteNotficationToken(token string) {
	var w = []string{token}

	var notficationPoolItem []model.NotficationPool

	err := initializers.DB.Where("token_array @> ?", pq.StringArray(w)).Find(&notficationPoolItem).Error
	if err != nil {
		loghandler.AppLogger.Error(err.Error())
	}
	fmt.Println(notficationPoolItem)

	for _, noti := range notficationPoolItem {
		noti = model.NotficationPool{
			UserId:     noti.UserId,
			TokenArray: deleteFromSliceByIndex(noti.TokenArray, slices.Index(noti.TokenArray, token)),
		}

		err := initializers.DB.Model(model.NotficationPool{}).Where("user_id = ?", noti.UserId).Updates(model.NotficationPool{UserId: noti.UserId, TokenArray: noti.TokenArray}).Error
		if err != nil {
			loghandler.AppLogger.Error(err.Error())
		}
	}
}

func deleteFromSliceByIndex(arr []string, index int) []string {
	if index < 0 || index >= len(arr) {
		return nil // Handle invalid index (optional)
	}
	return removeDuplicates(append(arr[:index], arr[index+1:]...))
}

func removeDuplicates(slice []string) []string {
	unique := make(map[string]bool)
	result := []string{}
	for _, v := range slice {
		if _, ok := unique[v]; !ok {
			unique[v] = true
			result = append(result, v)
		}
	}
	return result
}
