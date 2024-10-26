package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

func Config() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	viper.AddConfigPath(filepath.Join(basePath)) // 설정 파일 경로를 config 폴더로 설정

	// viper 설정
	viper.SetConfigName("config") // 설정 파일 이름 (확장자 제외)
	viper.SetConfigType("yaml")   // 설정 파일 타입
	//viper.AddConfigPath(".")      // 설정 파일 경로

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config 파일을 불러오는 곳에서 오류가 발생하였습니다., %s\n", err)
	}
}

func GetToken() string {
	return viper.GetString("discord.token")
}

func GetMongoConfig() map[string]interface{} {
	config := viper.GetStringMap("mongo")
	log.Println("MongoDB 설정 확인:", config) // Mongo 설정 값 출력
	return config
}
