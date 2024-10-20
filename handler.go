package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == s.State.User.ID {
		return
	}

	//if message.Content == "/schedule" {
	//	// 유저 정보 가져오기
	//	_, err := s.GuildMember(message.GuildID, message.Author.ID)
	//	if err != nil {
	//		s.ChannelMessageSend(message.ChannelID, "유저 정보를 가져올 수 없습니다.")
	//		return
	//	}
	//
	//	// 해당 채널에서 사용자의 권한 확인
	//	userPermissions, err := s.UserChannelPermissions(message.Author.ID, message.ChannelID)
	//	if err != nil {
	//		s.ChannelMessageSend(message.ChannelID, "사용자 권한을 확인할 수 없습니다.")
	//		return
	//	}
	//
	//	// 관리자인지 체크
	//	if userPermissions&discordgo.PermissionAdministrator != 0 {
	//		// 관리자인 경우 스케줄러 실행
	//		scheduler.Scheduler(s, message.ChannelID)
	//		s.ChannelMessageSend(message.ChannelID, "스케줄러가 설정되었습니다.")
	//	} else {
	//		// 관리자가 아닌 경우 경고 메시지 전송
	//		s.ChannelMessageSend(message.ChannelID, "해당 명령어는 관리자만 사용할 수 있습니다. 관리자 계정으로 다시 시도하세요.")
	//	}
	//}

	if message.Content == "/test" {
		s.ChannelMessageSend(message.ChannelID, "Test is succesful. Please do another test")
	}

	if message.Content == "/help" {
		s.ChannelMessageSend(message.ChannelID,
			"# Study-Bot (beta 0.1v) 명령어 🤖 \n "+
				"```"+
				"/test : 테스트 명령어 입니다. \n"+
				"/add @User : 사용자를 등록할 수 있습니다. \n"+
				"/find @User : 사용자가입 여부를 확인할 수 있습니다. \n"+
				"\n"+
				"```")
	}

	if strings.HasPrefix(message.Content, "/add") {
		if len(message.Mentions) > 0 {
			for _, mention := range message.Mentions {
				if !existUserByUserName(mention.Username, GetMongoConfig()) {
					saveMentionedUser(mention.ID, mention.Username, GetMongoConfig())
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

	if strings.HasPrefix(message.Content, "/find") {
		if len(message.Mentions) > 0 {
			for _, mention := range message.Mentions {
				if !existUserByUserName(mention.Username, GetMongoConfig()) {
					s.ChannelMessageSend(message.ChannelID,
						"```"+
							mention.Username+"님은 존재하지 않습니다. \n"+
							"/help 명령어를 통해 자세한 방법을 확인해주세요 ⚙️"+
							"```")
				} else {
					s.ChannelMessageSend(message.ChannelID,
						"```"+
							mention.Username+"님은 가입된 유저입니다. 🏃"+
							"```")
				}
			}
		}
	}
}
