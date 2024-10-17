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
				"/test : 테스트 명령어 입니다. \n" +
				"/add @User : 사용자를 등록할 수 있습니다. \n"+
				"\n"+
				"```")
	}

	if strings.HasPrefix(message.Content, "/add") {
		if len(message.Mentions) > 0 {
			for _, mention := range message.Mentions {
				if !existUserByUserName(mention.Username, getMongoConfig()) {
					saveMentionedUser(mention.ID, mention.Username, getMongoConfig())
					s.ChannelMessageSend(message.ChannelID,
						"```"+
							mention.Username+"님이 추가되었습니다!\n"+
							"```")
					continue
				}
				s.ChannelMessageSend(message.ChannelID,
					"```"+
						mention.Username+"은 존재하는 유저입니다!\n"+
						"```")
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

	if string.HasPrefix(message.Content, "/find") {
		if len(message.Mentions) > 0 {
			for _, mention := range message.Mentions {
				if !existUserByUserName(mention.Username, getMongoConfig()) {
					s.ChannelMessageSend(message.ChannelID,
						"```" +
							mention.Username + "님은 존재하지 않습니다. \n" +
							"/help 명령어를 통해 자세한 방법을 확인해주세요 ⚙️"
						"```")
				} else {
					s.ChannelMessageSend(message.ChannelID,
						"```" +
							mention.Username + "님은 가입된 유저입니다. 🏃"
						"```")
				}
			}
		}
}
