package handler

const (
	MentionCautionMassage = "```" +
		"해당 명령어는 멘션이 필요합니다 \n" +
		"사용 예시\n" +
		"/add @User" +
		"```"
	HelpMessage = "# Study-Bot (beta 0.1v) 명령어 🤖 \n " +
		"```" +
		"/test : 테스트 명령어 입니다. \n" +
		"/add @User : 사용자를 등록할 수 있습니다. \n" +
		"/find @User : 사용자가입 여부를 확인할 수 있습니다. \n" +
		"\n" +
		"```"
	AdminCautionMessage = "해당 명령어는 관리자만 사용할 수 있습니다. 관리자 계정으로 다시 시도하세요."
)

func getUserNotExistMessage(userName string) string {
	return "```" +
		userName + "님은 존재하지 않습니다. \n" +
		"/help 명령어를 통해 자세한 방법을 확인해주세요 ⚙️" +
		"```"
}

func getUserAlreadyExistMessage(userName string) string {
	return "```" +
		userName + "님은 가입된 유저입니다. 🏃" +
		"```"
}
