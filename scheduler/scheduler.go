package scheduler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

func Scheduler(session *discordgo.Session, channelID string) {
	currentTime := cron.New()

	// 매일 22:30에 메시지 전송
	_, err := currentTime.AddFunc("30 22 * * *", func() {
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
