package main

import (
	"fmt"
	"os"
	"os/signal"
	"study-bot-go/config"
	"study-bot-go/listener"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// 설정 초기화
	config.LoadConfig()
	//repository.Init()

	token := config.GetDiscordToken()
	dbConfig := config.GetMongoConfig()

	// MongoDB 설정 출력 확인
	fmt.Println("Loaded MongoDB LoadConfig:", dbConfig)

	// MongoDB 연결 시도
	if dbConfig.URI != "" {
		ConnectMongoDb(dbConfig)
	} else {
		fmt.Println("MongoDB 설정이 없습니다. config.yaml 파일을 확인하세요.")
		return
	}

	// 환경 변수에서 토큰을 가져옵니다 (또는 직접 토큰을 입력할 수도 있습니다)
	if token == "" || dbConfig.URI == "" {
		fmt.Println("No token provided. Please set DISCORD_BOT_TOKEN or DB environment variable.")
		return
	}

	// 새로운 디스코드 세션을 시작합니다
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// 메시지 생성 이벤트 핸들러를 등록합니다
	dg.AddHandler(listener.MessageListener)

	// WebSocket을 통해 디스코드와의 연결을 엽니다
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// 프로그램이 종료될 때까지 대기합니다
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// 디스코드 세션을 종료합니다
	dg.Close()
}
