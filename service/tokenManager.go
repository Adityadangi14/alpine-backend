package service

import (
	initializers "project_mine/initlizers"
	"project_mine/model"
)

func HandleTokenForUser(userId string, token string) (bool, error) {
	var notficationPoolItem model.NotficationPool

	rows := initializers.DB.Where("user_id = ?", userId).Find(&notficationPoolItem).RowsAffected

	if rows == 0 {
		var notficationPoolItem = model.NotficationPool{UserId: userId, TokenArray: []string{token}}
		err := initializers.DB.Create(&notficationPoolItem).Error

		if err != nil {
			return false, err
		}

		return true, nil
	} else {

		notficationPoolItem.TokenArray = append(notficationPoolItem.TokenArray, token)

		err := initializers.DB.Save(&notficationPoolItem).Error

		if err != nil {
			return false, err
		}

		return true, nil

	}

}
