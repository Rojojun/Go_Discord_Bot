package exception

import "log"

func MessageSendFailureException(sendErr error) {
	log.Fatal("메시지 전송 실패:", sendErr)
}

func UserAuthorizationException(sendErr error) {
	log.Fatal("사용자 권한을 확인할 수 없습니다. : ", sendErr)
}
