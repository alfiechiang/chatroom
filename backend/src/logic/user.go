package logic

import (
	"context"
	"fmt"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	MessageChannel chan *Message `json:"-"`
	Nickname       string
	EnterAt        time.Time `json:"enter_at"`

	Conn *websocket.Conn
}

func (u *User) SendMessage(ctx context.Context) {
	for msg := range u.MessageChannel {
		fmt.Println(msg)
		wsjson.Write(ctx, u.Conn, msg)
	}

}
