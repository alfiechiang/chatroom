package logic

import (
	"context"
	"time"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	MessageChannel chan *Message `json:"-"`
	Nickname       string
	EnterAt        time.Time `json:"enter_at"`
	Ctx  context.Context
	Conn *websocket.Conn
}

func NewUser(nickname string,conn *websocket.Conn,ctx context.Context)(*User){
	return &User{
		MessageChannel: make(chan *Message),
		Nickname:       nickname,
		EnterAt:        time.Now(),
		Ctx: ctx,
		Conn:           conn,
	}


}

func (u *User) SendMessage(ctx context.Context) {
	for msg := range u.MessageChannel {
		wsjson.Write(ctx, u.Conn, msg)
	}

}
