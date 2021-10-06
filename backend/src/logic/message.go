package logic

type Message struct {
	// 哪个用户发送的消息
	//	User    *User     `json:"user"`
	Type    int       `json:"type"`
	Content string `json:"content"`
	// MsgTime time.Time `json:"msg_time"`

	// ClientSendTime time.Time `json:"client_send_time"`

	// 消息 @ 了谁
	Ats []string `json:"ats"`

	// 用户列表不通过 WebSocket 下发
	// Users []*User `json:"users"`
}