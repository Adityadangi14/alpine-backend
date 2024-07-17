package service

import (
	"fmt"
	initializers "project_mine/initlizers"
	"project_mine/model"
	"slices"

	"github.com/lib/pq"
)

func DeleteNotficationToken(token string) {
	var w = []string{token}

	var notficationPoolItem model.NotficationPool

	initializers.DB.Where("token_array @> ?", pq.StringArray(w)).Find(&notficationPoolItem)

	modifiedArray := deleteFromSliceByIndex(notficationPoolItem.TokenArray, slices.Index(notficationPoolItem.TokenArray, token))

	notficationPoolItem = model.NotficationPool{
		UserId:     notficationPoolItem.UserId,
		TokenArray: modifiedArray,
	}

	err := initializers.DB.Save(&notficationPoolItem).Error

	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func deleteFromSliceByIndex(arr []string, index int) []string {
	if index < 0 || index >= len(arr) {
		return nil // Handle invalid index (optional)
	}
	return append(arr[:index], arr[index+1:]...)
}
