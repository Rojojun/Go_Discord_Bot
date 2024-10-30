package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"study-bot-go/exception"
	"study-bot-go/repository"
	"study-bot-go/scheduler"
)

func FindDiscordUser(session *discordgo.Session, message *discordgo.MessageCreate) {
	member, err := session.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "유저 정보를 가져올 수 없습니다."); sendErr != nil {
			exception.MessageSendFailureException(sendErr)
		}
		return
	}

	log.Printf("유저 정보: %+v\n", member)
}

func DoActionWithPermission(session *discordgo.Session, message *discordgo.MessageCreate) {
	// 서버의 소유자 ID 가져오기
	guild, err := session.State.Guild(message.GuildID)
	if err != nil {
		fmt.Println("Error retrieving guild:", err)
		return
	}

	userID := message.Author.ID
	userPermission := findPermission(session, message)

	// 메시지를 보낸 유저가 서버 오너인지 확인
	if userID == guild.OwnerID || userPermission&discordgo.PermissionAdministrator == 0 || userPermission&discordgo.PermissionAdministrator == 8 {
		scheduler.Scheduler(session, message.ChannelID)
		sendSchedulerSetSuccessMassage(session, message)
	} else {
		sendSchedulerSetFailedMassage(session, message)
	}
}

func GetResignMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	repository.DeleteUserByUserName(message.Author.Username, message.GuildID)
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, getResignMessage(message.Author.Username)); sendErr != nil {
		log.Fatalln("/bye 명령어 전송 실패 >>>> ", sendErr)
	}
}

func GetDailyGoalSettingMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	user, err := repository.FindUserBy(message.Author.Username, message.GuildID)
	if err != nil {
		fmt.Println("User 조회 오류:", err)
		return
	}

	if user == nil {
		fmt.Println("해당 유저가 존재하지 않습니다.")
		return
	}

	goal, err := repository.FindDailyGoalByOwnerId(user.Id)
	if err != nil {
		fmt.Println("Goal 조회 오류:", err)
		return
	}

	if goal != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "이미 목표가 등록되어 있습니다.\n`/find goal daily` 명령어로 목표를 확인해주세요."); sendErr != nil {
			log.Fatalln("/goal daily 명령어 전송 실패 >>>> ", sendErr)
		}
		return
	}

	dailyGoal := strings.TrimPrefix(message.Content, "/goal daily ")

	if _, sendErr := session.ChannelMessageSend(message.ChannelID, getAddDailyGoalMessage(dailyGoal)); sendErr != nil {
		log.Fatalln("/goal daily 명령어 전송 실패 >>>> ", sendErr)
	}

	repository.SaveGoal(dailyGoal, user.Id, "DAILY")
}

func GetHelpMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, HelpMessage); sendErr != nil {
		log.Fatalln("/help 명령어 전송 실패 >>>> ", sendErr)
	}
}

func GetTestMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, "테스트 명령어 성공"); sendErr != nil {
		exception.UnstableServerConnectionException(sendErr)
	}
}

func AddUserWithMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if len(message.Mentions) > 0 {
		addMentionedUsers(session, message, message.Mentions)
	} else {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, MentionCautionMassage); sendErr != nil {
			exception.UnstableServerConnectionException(sendErr)
		}
	}
}

func ExistUserByMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if len(message.Mentions) > 0 {
		existAddUsers(session, message, message.Mentions)
	}
}

func GetUserHasGoal(session *discordgo.Session, message *discordgo.MessageCreate) {
	user, err := repository.FindUserBy(message.Author.Username, message.GuildID)
	if err != nil {
		fmt.Println("User 조회 오류:", err)
		return
	}

	if user == nil {
		fmt.Println("해당 유저가 존재하지 않습니다.")
		return
	}

	goal, err := repository.FindDailyGoalByOwnerId(user.Id)
	if err != nil {
		fmt.Println("Goal 조회 오류:", err)
		return
	}

	if goal != nil {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "나의 일간 목표 : "+goal.Goal); sendErr != nil {
			log.Fatalln("/find goal daily 명령어 전송 실패 >>>> ", sendErr)
		}
		return
	} else {
		if _, sendErr := session.ChannelMessageSend(message.ChannelID, "등록된 목표가 없습니다.\n`/goal daily 목표` 명령어로 목표를 등록해주세요."); sendErr != nil {
			log.Fatalln("/find goal daily 명령어 전송 실패 >>>> ", sendErr)
		}
	}
}

func _____private__area_____() {}

func addMentionedUsers(session *discordgo.Session, message *discordgo.MessageCreate, users []*discordgo.User) {
	for _, mention := range users {
		if !repository.ExistUserByUserName(mention.Username, message.GuildID) {
			repository.SaveMentionedUser(mention.ID, mention.Username, message.GuildID)
			sendUserAddSuccessMessage(session, message, mention)
			continue
		}
		sendMessageOf("```"+
			mention.Username+"은 존재하는 유저입니다!\n"+
			"```", session, message)
	}
}

func existAddUsers(session *discordgo.Session, message *discordgo.MessageCreate, users []*discordgo.User) {
	for _, userInfo := range users {
		if !repository.ExistUserByUserName(userInfo.Username, message.GuildID) {
			sendMessageOf(getUserNotExistMessage(userInfo.Username), session, message)
		} else {
			sendMessageOf(getUserAlreadyExistMessage(userInfo.Username), session, message)
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
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, AdminCautionMessage); sendErr != nil {
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

func sendMessageOf(content string, session *discordgo.Session, message *discordgo.MessageCreate) {
	if _, sendErr := session.ChannelMessageSend(message.ChannelID, content); sendErr != nil {
		exception.MessageSendFailureException(sendErr)
	}
}
