package server

import (
	"fmt"
	"learn-one/cmd/common/interfaces/user"
)

type EmailService struct {
	// 可以依赖 user 包
	UserService *user.UserService
}

func (e *EmailService) SendNotification(message string) {
	// 实现接口
	fmt.Println("发送邮件:", message)
}
