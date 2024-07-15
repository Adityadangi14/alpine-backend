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

	response, err := fcmClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: Title,
			Body:  Body,
		},
		Token: "d6kB-iqjR2upwy4La-OZfs:APA91bFUy_1hCNZXXZm2WhyvJbQb0V47GksjPgc_8C7L8dk9E9e1N2J-6mrKxHsJ9dE8WQxUy5PkCMiOQTFGQq8Vb09qPS86x7G6lbWSFW1Mfr5jUpf-kV_5WUVBNYtW6TUYQv_omEgG", // it's a single device token
	})

	if err != nil {
		fmt.Printf("Error in initializing firebase app: %s", err)
	}

	fmt.Println(response)
}
