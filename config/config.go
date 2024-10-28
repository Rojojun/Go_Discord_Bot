package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Mongo   MongoConfig   `mapstructure:"mongo"`
	Discord DiscordConfig `mapstructure:"discord"`
}

type MongoConfig struct {
	URI        string `mapstructure:"uri"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

type DiscordConfig struct {
	Token string `mapstructure:"token"`
}

var AppConfig Config

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	configPath := filepath.Join(basePath, "..", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("[FATAL] 설정 파일이 존재하지 않습니다: %s", configPath)
	} else {
		fmt.Printf("[SUCCESS] 설정 파일이 존재합니다: %s\n", configPath)
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

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("설정 파일을 구조체에 매핑하는 중 오류 발생: %s\n", err)
	}
}

func GetMongoConfig() MongoConfig {
	return AppConfig.Mongo
}

func GetDiscordToken() string {
	return AppConfig.Discord.Token
}
