package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == s.State.User.ID {
		return
	}

	if message.Content == "/test" {
		s.ChannelMessageSend(message.ChannelID, "Test is succesful. Please do another test")
	}

	if message.Content == "/help" {
		s.ChannelMessageSend(message.ChannelID,
			"# Study-Bot (beta 0.1v) 명령어 🤖 \n "+
				"```"+
				"/test : 테스트 명령어 입니다. \n"+
				"\n"+
				"```")
	}

	if strings.HasPrefix(message.Content, "/add") {
		if len(message.Mentions) > 0 {
			for _, mention := range message.Mentions {
				findUserByUserName()
			}
		} else {
			s.ChannelMessageSend(message.ChannelID,
				"```"+
					"해당 명령어는 멘션이 필요합니다 \n"+
					"사용 예시\n"+
					"/add @User"+
					"```")
		}
	}
}
