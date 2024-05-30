package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func Config() {
	// viper 설정
	viper.SetConfigName("config") // 설정 파일 이름 (확장자 제외)
	viper.SetConfigType("yaml")   // 설정 파일 타입
	viper.AddConfigPath(".")      // 설정 파일 경로

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config 파일을 불러오는 곳에서 오류가 발생하였습니다., %s\n", err)
	}
}

func GetToken() string {
	return viper.GetString("discord.token")
}