package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// 설정 초기화
	Config()

	// 설정 파일에서 토큰을 가져옵니다
	token := GetToken()

    // 환경 변수에서 토큰을 가져옵니다 (또는 직접 토큰을 입력할 수도 있습니다)
    if token == "" {
        fmt.Println("No token provided. Please set DISCORD_BOT_TOKEN environment variable.")
        return
    }

    // 새로운 디스코드 세션을 시작합니다
    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Println("Error creating Discord session,", err)
        return
    }

    // 메시지 생성 이벤트 핸들러를 등록합니다
    dg.AddHandler(messageCreate)

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

// 메시지가 생성될 때 호출되는 핸들러 함수입니다
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // 봇이 보낸 메시지는 무시합니다
    if m.Author.ID == s.State.User.ID {
        return
    }

    if m.Content == "!test" {
        s.ChannelMessageSend(m.ChannelID, "Test is succesful. Please do another test")
    }  	
	if m.Content == "!help" {
		s.ChannelMessageSend(m.ChannelID,
		"# 군필소녀 디스코드 봇 사용방법 🤖 \n " +	
		"```" +
		"!test : 테스트 명령어 입니다. \n" +
		"\n" + 
		"!test : 테스트 명령어 입니다. \n" +
		"```")
	}
}