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
	log.Println("MongoDB Connection Success")

	uri := fmt.Sprintf("%s", mongoConfig["uri"].(string))

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB 클라이언트 생성 실패:", err)
	}

	fmt.Println("MongoDB에 연결되었습니다.")
}

func existUserByUserName(userName string, mongoConfig map[string]interface{}) bool {
	collection := client.Database(mongoConfig["database"].(string)).Collection(mongoConfig["collection"].(string))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userName": userName}

	err := collection.FindOne(ctx, filter).Err()
	if err == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func saveMentionedUser(userId, userName string, mongoConfig map[string]interface{}) {
	collection := client.Database(mongoConfig["database"].(string)).Collection(mongoConfig["collection"].(string))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, map[string]interface{}{
		"userId":    userId,
		"userName":  userName,
		"createdAt": time.Now(),
	})

	if err != nil {
		log.Println("사용자 저장 실패 : ", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다 : ", userName)
	}
}
