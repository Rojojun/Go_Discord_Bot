package scheduler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"time"
)

func Scheduler(session *discordgo.Session, channelID string) {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("지역 설정에 오류가 발생하였습니다.", err)
		return
	}
	currentTime := cron.New(cron.WithLocation(location))

	// 매일 22:30에 메시지 전송
	_, err = currentTime.AddFunc("50 23 * * *", func() {
		runAlert(session, channelID)
	})
	if err != nil {
		fmt.Println("Error scheduling the task:", err)
		return
	}

	// 스케줄러 시작
	currentTime.Start()
	fmt.Println("Scheduler started. Waiting for the task to run...")
}

func runAlert(session *discordgo.Session, channelID string) {
	_, err := session.ChannelMessageSend(channelID, "10시 30분!")
	if err != nil {
		fmt.Println("Error sending the message:", err)
	}
}
