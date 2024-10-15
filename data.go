package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

func ConnectMongoDb(mongoConfig map[string]interface{}) {

	uri := fmt.Sprintf("mongodb://%s/%s",
		mongoConfig["uri"],
		mongoConfig["database"])

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB 클라이언트 생성 실패:", err)
	}

	fmt.Println("MongoDB에 연결되었습니다.")
}

func findUserByUserName(userName string, mongoConfig map[string]string) {
	collection := client.Database(mongoConfig["database"]).Collection(mongoConfig["collection"])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userName": userName}

	_, err := collection.FindOne(ctx)
}

func saveMentionedUser(userName string, mongoConfig map[string]string) {
	collection := client.Database(mongoConfig["database"]).Collection(mongoConfig["collection"])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, map[string]interface{}{
		"userName":  userName,
		"createdAt": time.Now(),
	})

	if err != nil {
		log.Println("사용자 저장 실패 : ", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다 : ", userName)
	}
}
