package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetRandomWord() {
	db := getConnection()
	col := db.Collection("word")

	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 1}}}}, bson.D{{"$regex", bson.D{{"$regex", "^[\\s\\S]{40,}$"}}}}}
	word, err := col.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word)

}

func getConnection() *mongo.Database {
	type User struct {
		name string `bson:"name"`
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/lingoDB"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("lingoDB")
	return db
	//col := db.Collection("users")

	//var users bson.M
	//if err = col.FindOne(ctx, bson.M{"user":"Henk"}).Decode(&users); err != nil {
	//	log.Fatal(err)
	//}
	//
	//return users["user"], err

}
