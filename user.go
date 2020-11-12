package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func initializeApp() *firebase.App {
	sa := option.WithCredentialsFile("configs/lingoflutterapp-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		log.Fatalln(err)
	}
	return app
}
func getData(app firebase.App) *firestore.DocumentSnapshot {
	client, err := app.Firestore(context.Background())
	result, err := client.Collection("test").Doc("testDoc").Get(context.Background())

	if err != nil {

	} else {
		return result
	}
	return result
}
