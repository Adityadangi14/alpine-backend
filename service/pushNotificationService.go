package service

import (
	"context"
	"fmt"
	loghandler "project_mine/logHandler"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func PushNotificationService(Body string, Title string) {

	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		fmt.Printf("Error in initializing firebase app: %s", err)

	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		fmt.Printf("Error in initializing firebase app: %s", err)
	}
	tokens, err := GetTokensForUserNotification()

	fmt.Println("tokens", len(tokens))

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		fmt.Printf("Error in getting tokens: %s", err)
	}
	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{

		Notification: &messaging.Notification{
			Title: Title,
			Body:  Body,
		},

		Tokens: tokens,
	})

	if err != nil {
		loghandler.AppLogger.Error(string(err.Error()))
		fmt.Printf("Error in initializing firebase app: %s", err)
	}

	for index, res := range response.Responses {
		if res.Error != nil {
			loghandler.AppLogger.Error(string(res.Error.Error()))
			fmt.Println(index, res.Error)
			tkn := tokens[index]
			DeleteNotficationToken(tkn)
		}
	}

}
