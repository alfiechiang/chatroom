package logic

import (
	"fmt"
	"time"
)

type Message struct {
	// 哪个用户发送的消息
	//	User    *User     `json:"user"`
	Type    int    `json:"type"`
	Content string `json:"content"`
	// MsgTime time.Time `json:"msg_time"`

	// ClientSendTime time.Time `json:"client_send_time"`

	// 消息 @ 了谁
	Ats      []string `json:"ats"`
	Nickname string   `json:"nickname"`

	// 用户列表不通过 WebSocket 下发
	// Users []*User `json:"users"`
}

func WelcomeMessage(nickname string) *Message {
	fmt.Println("WelcomeMessage")
	return &Message{
		Content: nickname + "您好!!，欢迎加入聊天室！",
		Type:    0,
	}

}

func InviteMessage(nickname string) *Message {

	t := time.Now()

	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return &Message{
		Content: nickname + "加入时间：" + formatted,
		Type:    2,
	}
}

func ChatMessage(content string, ty int, nickname string) *Message {
	return &Message{
		Content:  content,
		Type:     ty,
		Nickname: nickname,
	}
}
