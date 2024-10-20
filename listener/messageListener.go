package listener

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"study-bot-go/handler"
	"time"
)

func MessageListener(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "/help" {
		handler.GetHelpMessage(session, message)
		log.Println("/help 메시지 전송 성공 : ", time.Now())
	}

	if message.Content == "/schedule" {
		handler.FindDiscordUser(session, message)
		handler.DoActionWithPermission(session, message)
		log.Printf("/schedule 메시지 전송 성공 : ", time.Now())
	}

	if message.Content == "/test" {
		handler.GetTestMessage(session, message)
	}

	if strings.HasPrefix(message.Content, "/add") {
		handler.AddUserWithMessage(session, message)
	}

	if strings.HasPrefix(message.Content, "/find") {

	}
}
