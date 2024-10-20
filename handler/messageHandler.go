package handler

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"study-bot-go/exception"
	"study-bot-go/repository"
	"study-bot-go/scheduler"
)

const MENTION_CAUTION_MASSAGE = "```" +
	"해당 명령어는 멘션이 필요합니다 \n" +
	"사용 예시\n" +
	"/add @User" +
	"```"

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

func GetHelpMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, ""+
		"# Study-Bot (beta 0.1v) 명령어 🤖 \n "+
		"```"+
		"/test : 테스트 명령어 입니다. \n"+
		"/add @User : 사용자를 등록할 수 있습니다. \n"+
		"/find @User : 사용자가입 여부를 확인할 수 있습니다. \n"+
		"\n"+
		"```"); sendErr != nil {
		log.Fatalln("/help 명령어 전송 실패 >>>> ", sendErr)
	}
}

func GetTestMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "테스트 명령어 성공"); sendErr != nil {
		log.Fatalln("서버 연결을 확인해주세요.")
	}
}

func AddUserWithMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if len(message.Mentions) > 0 {
		addMentionedUsers(session, message, message.Mentions)
	} else {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, MENTION_CAUTION_MASSAGE); sendErr != nil {
			log.Fatalln()
		}
	}
}

func _____private__area_____() {}

func addMentionedUsers(session *discordgo.Session, message *discordgo.MessageCreate, users []*discordgo.User) {
	for _, mention := range users {
		if !repository.ExistUserByUserName(mention.Username) {
			repository.SaveMentionedUser(mention.ID, mention.Username)
			sendUserAddSuccessMessage(session, message, mention)
			continue
		}
	}
}

func findPermission(session *discordgo.Session, message *discordgo.MessageCreate) int64 {
	userPermissions, err := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "사용자 권한을 확인할 수 없습니다."); sendErr != nil {
			exception.UserAuthorizationException(sendErr)
		}
		return 0
	}
	return userPermissions
}

func sendSchedulerSetSuccessMassage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "스케줄러가 설정되었습니다."); sendErr != nil {
		exception.MessageSendFailureException(sendErr)
	}
}

func sendSchedulerSetFailedMassage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "해당 명령어는 관리자만 사용할 수 있습니다. 관리자 계정으로 다시 시도하세요."); sendErr != nil {
		exception.MessageSendFailureException(sendErr)
	}
}

func sendUserAddSuccessMessage(session *discordgo.Session, message *discordgo.MessageCreate, user *discordgo.User) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID,
		"```"+
			user.Username+"님이 추가되었습니다!\n"+
			"```"); sendErr != nil {
		exception.MessageSendFailureException(sendErr)
	}
}

func sendMessageOf(userExistMessage string, session *discordgo.Session, message *discordgo.MessageCreate) {

}
