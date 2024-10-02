package main

import (
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/test" {
		s.ChannelMessageSend(m.ChannelID, "Test is succesful. Please do another test")
	}
	if m.Content == "/help" {
		s.ChannelMessageSend(m.ChannelID,
			"# Study-Bot (beta 0.1v) 명령어 🤖 \n "+
				"```"+
				"/test : 테스트 명령어 입니다. \n"+
				"\n"+
				"```")
	}
}
