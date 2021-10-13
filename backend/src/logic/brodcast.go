package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// import "fmt"

type Brodcast struct {
	MessageChannel chan *Message
	Users          []*User
	Conn           *websocket.Conn
	Ctx            context.Context
}

func NewBrodcast() *Brodcast {
	return &Brodcast{
		MessageChannel: make(chan *Message),
		Users:          make([]*User, 0),
	}
}

func (b *Brodcast) Start() {

	for {
		select {
		case msg := <-b.MessageChannel:
			// 新用户进入
			for _, user := range b.Users {
				fmt.Println(user.Nickname)
				user.MessageChannel <- msg
			}
		}
	}

}

func (b *Brodcast) ReceiveMessage(u *User) {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {

		err = wsjson.Read(u.Ctx, u.Conn, &receiveMsg)
		if err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				fmt.Println("Websocket closing")
				return
			}
			log.Println(err)
			return
		}
		fmt.Println("ceffefe")

		ty, err := strconv.Atoi(receiveMsg["type"])
		if err != nil {
			fmt.Println("switch string to int faild")
		}

		msg := ChatMessage(receiveMsg["content"], ty, u.Nickname)
		b.MessageChannel <- msg
	}

}
