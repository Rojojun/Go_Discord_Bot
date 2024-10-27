package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func Config() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	configPath := filepath.Join(basePath, "..", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("설정 파일이 존재하지 않습니다: %s", configPath)
	} else {
		fmt.Printf("설정 파일이 존재합니다: %s\n", configPath)
	}

	// Viper 설정
	viper.AddConfigPath(filepath.Join(basePath, ".."))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config 파일을 불러오는 중 오류 발생: %s\n", err)
	} else {
		fmt.Println("Config 파일이 성공적으로 로드되었습니다.")
	}
}

func GetToken() string {
	return viper.GetString("discord.token")
}

func GetMongoConfig() map[string]interface{} {
	config := viper.GetStringMap("mongo")
	log.Println("MongoDB 설정 확인:", config) // Mongo 설정 값 출력
	if len(config) == 0 {
		log.Panic("MongoDB 설정이 잘못되었습니다. --- Config")
	}

	return config
}
