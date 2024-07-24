package service

import (
	"fmt"
	initializers "project_mine/initlizers"
	loghandler "project_mine/logHandler"
	"project_mine/model"
	"slices"
)

func HandleTokenForUser(userId string, token string) (bool, error) {
	var notficationPoolItem model.NotficationPool

	rows := initializers.DB.Where("user_id = ?", userId).Find(&notficationPoolItem).RowsAffected

	if rows == 0 {
		var notficationPoolItem = model.NotficationPool{UserId: userId, TokenArray: []string{token}}
		err := initializers.DB.Create(&notficationPoolItem).Error

		if err != nil {
			loghandler.AppLogger.Error(string(err.Error()))
			return false, err
		}

		return true, nil
	} else {

		if slices.Contains(notficationPoolItem.TokenArray, token) {
			return false, fmt.Errorf("token alredy present")
		}

		notficationPoolItem.TokenArray = append(notficationPoolItem.TokenArray, token)

		err := initializers.DB.Save(&notficationPoolItem).Error

		if err != nil {
			loghandler.AppLogger.Error(string(err.Error()))
			return false, err
		}

		return true, nil
	}
}
