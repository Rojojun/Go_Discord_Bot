package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if m.Content == "!log" {
		multilineLog := "User: " + m.Author.Username + "\n" +
			"Message: " + m.Content + "\n" +
			"Channel: " + m.ChannelID + "\n" +
			"Here is a line with backticks: `` ` ``."

		data := []interface{}{multilineLog}
		err := writeToSheet(data)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed to write to Google Sheets: "+err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, "Logged to Google Sheets succe