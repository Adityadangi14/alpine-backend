package service

import (
	initializers "project_mine/initlizers"
	"project_mine/model"
)

func GetTokensForUserNotification() ([]string, error) {

	//need to optimise
	var notficationPoolItems []model.NotficationPool

	err := initializers.DB.Find(&notficationPoolItems).Error

	if err != nil {
		return nil, err
	}

	var tokens []string

	for _, noti := range notficationPoolItems {
		tokens = append(tokens, noti.TokenArray...)
	}

	return tokens, nil

}
