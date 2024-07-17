package service

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func PushNotificationService(Body string, Title string) {

	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		fmt.Printf("Error in initializing firebase app: %s", err)

	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		fmt.Printf("Error in initializing firebase app: %s", err)
	}
	tokens, err := GetTokensForUserNotification()

	fmt.Println("tokens", len(tokens))

	if err != nil {
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
		fmt.Printf("Error in initializing firebase app: %s", err)
	}

	for index, resposne := range response.Responses {
		if resposne.Error != nil {
			fmt.Println(index, resposne.Error)
		}
	}

}
