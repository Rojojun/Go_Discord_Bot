package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"study-bot-go/config"
	"time"
)

var (
	mongoConnection config.MongoConfig
	retryCount      context.Context
	client          *mongo.Client
)

func ExistUserByUserName(userName string) bool {
	println("client ::::: ", client)
	config.GetMongoConfig()
	collection := client.Database(mongoConnection.Database).Collection(mongoConnection.Collection)
	filter := bson.M{"userName": userName}

	err := collection.FindOne(retryCount, filter).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false
	}
	return true
}

func SaveMentionedUser(userId, userName string) {
	collection := client.Database(mongoConnection.Database).Collection(mongoConnection.Collection)
	result, err := collection.InsertOne(retryCount, map[string]interface{}{
		"userId":    userId,
		"userName":  userName,
		"createdAt": time.Now(),
	})
	if err != nil {
		log.Println("사용자 저장 실패 : ", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다 : ", result.InsertedID)
	}
}

func existUserByUserName() *mongo.Collection {
	return client.Database(mongoConnection.Database).Collection(mongoConnection.Collection)
}

func setRetryCount() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return ctx
}

// MongoDB 연결 함수
func connectMongoDB() *mongo.Client {
	uri := mongoConnection.URI // MongoDB URI 가져오기
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(retryCount, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB 연결 실패: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(retryCount, nil); err != nil {
		log.Fatalf("MongoDB Ping 실패: %v", err)
	}

	fmt.Println("MongoDB에 연결되었습니다.")
	return client
}
