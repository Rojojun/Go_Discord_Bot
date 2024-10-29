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
	client          *mongo.Client
)

// 유저 이름으로 유저 존재 여부 확인 함수
func ExistUserByUserName(userName string, guildId string) bool {
	client, err := connectMongoDB()
	if err != nil {
		log.Println("MongoDB 연결 오류:", err)
		return false
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().Collection)
	filter := bson.M{"userName": userName, "guildId": guildId}

	println(userName)

	err = collection.FindOne(context.Background(), filter).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}

// 유저 정보 저장 함수
func SaveMentionedUser(userId, userName string, guildId string) {
	client, err := connectMongoDB()
	if err != nil {
		log.Println("MongoDB 연결 오류:", err)
		return
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().Collection)
	_, err = collection.InsertOne(context.Background(), bson.M{
		"userId":    userId,
		"userName":  userName,
		"guildId":   guildId,
		"createdAt": time.Now(),
	})
	if err != nil {
		log.Println("사용자 저장 실패:", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다.")
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

// MongoDB 클라이언트를 반환하는 함수
func connectMongoDB() (*mongo.Client, error) {
	// MongoDB 설정
	uri := config.GetMongoConfig().URI
	clientOptions := options.Client().ApplyURI(uri)

	// 컨텍스트와 클라이언트 생성
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB 연결 실패: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB Ping 실패: %w", err)
	}

	fmt.Println("MongoDB에 연결되었습니다.")
	return client, nil
}
