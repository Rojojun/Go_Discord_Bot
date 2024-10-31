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
	"study-bot-go/domain"
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

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionUser)
	filter := bson.M{"userName": userName, "guildId": guildId}

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

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionUser)
	_, err = collection.InsertOne(context.Background(), bson.M{
		"userId":      userId,
		"userName":    userName,
		"guildId":     guildId,
		"createdAt":   time.Now(),
		"setSchedule": false,
	})
	if err != nil {
		log.Println("사용자 저장 실패:", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다.")
	}
}

func FindDailyGoalByOwnerId(ownerId string) (*domain.Goal, error) {
	client, err := connectMongoDB()
	if err != nil {
		log.Println("MongoDB 연결 오류:", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionGoal)
	filter := bson.M{"ownerId": ownerId}

	var goal domain.Goal
	err = collection.FindOne(context.Background(), filter).Decode(&goal)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("해당 사용자와 서버의 문서를 찾을 수 없습니다.")
			return nil, nil
		}
		log.Fatalln("MongoDB에서 문서 검색 오류:", err)
		return nil, err
	}

	return &goal, nil
}

func DeleteUserByUserName(userName string, guildId string) {
	client, err := connectMongoDB()
	if err != nil {
		log.Println("MongoDB 연결 오류:", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionUser)
	filter := bson.M{"userName": userName, "guildId": guildId}

	_, err = collection.DeleteOne(context.Background(), filter)
}

func SaveGoal(goal string, ownerId string, goalType string) {
	client, err := connectMongoDB()
	if err != nil {
		log.Println("MongoDB 연결 오류:", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionGoal)
	_, err = collection.InsertOne(context.Background(), bson.M{
		"goal":        goal,
		"ownerId":     ownerId,
		"goalType":    goalType,
		"createdAt":   time.Now(),
		"setSchedule": false,
	})
	if err != nil {
		log.Println("사용자 저장 실패:", err)
	} else {
		fmt.Println("사용자가 MongoDB에 저장되었습니다.")
	}
}

func FindUserBy(userName string, guildId string) (*domain.User, error) {
	client, err := connectMongoDB()
	if err != nil {
		log.Fatalln("MongoDB 연결 오류:", err)
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionUser)
	filter := bson.M{"userName": userName, "guildId": guildId}

	var user domain.User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("해당 사용자와 서버의 문서를 찾을 수 없습니다.")
			return nil, nil
		}
		log.Fatalln("MongoDB에서 문서 검색 오류:", err)
		return nil, err
	}

	return &user, nil
}

func SetSchedule(goal *domain.Goal, id string, s string) error {
	client, err := connectMongoDB()
	if err != nil {
		log.Fatalln("MongoDB 연결 오류:", err)
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		_ = client.Disconnect(ctx)
	}(client, context.Background())

	collection := client.Database(config.GetMongoConfig().Database).Collection(config.GetMongoConfig().CollectionUser)
	filter := bson.M{}
}

//func existUserByUserName() *mongo.Collection {
//	return client.Database(mongoConnection.Database).Collection(mongoConnection.Collection)
//}

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
