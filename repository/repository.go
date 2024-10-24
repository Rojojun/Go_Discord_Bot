package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"study-bot-go/config"
	"time"
)

var (
	mongoConnection = config.GetMongoConfig()
	collection      = existUserByUserName()
	retryCount      = setRetryCount()
	client          *mongo.Client
)

func init() {
	mongoConnection = config.GetMongoConfig()
	retryCount = setRetryCount()
	client = connectMongoDB()

	// mongoConnection["database"] 및 ["collection"]이 nil인지 확인 후 처리
	if mongoConnection["database"] == nil || mongoConnection["collection"] == nil {
		log.Panic("MongoDB 설정이 잘못되었습니다.")
	}

	collection = client.Database(mongoConnection["database"].(string)).Collection(mongoConnection["collection"].(string))
}

func ExistUserByUserName(userName string) bool {
	filter := bson.M{"userName": userName}

	err := collection.FindOne(retryCount, filter).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false
	}
	return true
}

func SaveMentionedUser(userId, userName string) {
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
	return client.Database(mongoConnection["database"].(string)).Collection(mongoConnection["collection"].(string))
}

func setRetryCount() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ctx
}

// MongoDB 연결 함수
func connectMongoDB() *mongo.Client {
	// MongoDB 연결 로직 추가
	// 클라이언트를 반환
	return client
}
