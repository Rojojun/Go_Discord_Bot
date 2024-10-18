package handler

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"study-bot-go/scheduler"
)

func FindDiscordUser(session *discordgo.Session, message *discordgo.MessageCreate) {
	member, err := session.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "유저 정보를 가져올 수 없습니다."); sendErr != nil {
			log.Fatalln("메시지 전송 실패:", sendErr)
		}
		return
	}

	log.Printf("유저 정보: %+v\n", member)
}

func DoActionWithPermission(session *discordgo.Session, message *discordgo.MessageCreate) {
	userPermission := findPermission(session, message)
	if userPermission&discordgo.PermissionAdministrator != 0 {
		scheduler.Scheduler(session, message.ChannelID)
		sendSchedulerSetSuccessMassage(session, message)
	} else {
		sendSchedulerSetFailedMassage(session, message)
	}
}

func _____private__area_____() {}

func findPermission(session *discordgo.Session, message *discordgo.MessageCreate) int64 {
	userPermissions, err := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "사용자 권한을 확인할 수 없습니다."); sendErr != nil {
			log.Fatal("사용자 권한을 확인할 수 없습니다.")
		}
		return 0
	}
	return userPermissions
}

func sendSchedulerSetSuccessMassage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "스케줄러가 설정되었습니다."); sendErr != nil {
		log.Fatalln("메시지 전송 실패:", sendErr)
	}
}

func sendSchedulerSetFailedMassage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "해당 명령어는 관리자만 사용할 수 있습니다. 관리자 계정으로 다시 시도하세요."); sendErr != nil {
		log.Fatalln("메시지 전송 실패:", sendErr)
	}
}
