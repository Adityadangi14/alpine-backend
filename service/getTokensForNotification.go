package service

import (
	initializers "project_mine/initlizers"
	loghandler "project_mine/logHandler"
	"project_mine/model"
)

func GetTokensForUserNotification() ([]string, error) {

	//need to optimise
	var notficationPoolItems []model.NotficationPool

	err := initializers.DB.Find(&notficationPoolItems).Error

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		return nil, err
	}

	var tokens []string

	for _, noti := range notficationPoolItems {
		tokens = append(tokens, noti.TokenArray...)
	}

	return tokens, nil

}
