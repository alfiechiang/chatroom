package controller

import (
	"backend/src/logic"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

var Users []*logic.User = make([]*logic.User, 0)

func ChatWebSocketInit(c *gin.Context) {

	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return
	}

	nickname := c.Request.FormValue("nickname")
	//defer conn.Close(websocket.StatusInternalError, "Internal Error")

	_, cancel := context.WithTimeout(c, time.Second*30)
	defer cancel()

	user := &logic.User{
		MessageChannel: make(chan *logic.Message),
		Nickname:       nickname,
		EnterAt:        time.Now(),
		Conn:           conn,
	}

	go user.SendMessage(c)

	user.MessageChannel <- &logic.Message{
		Content: nickname + "您好，欢迎加入聊天室！",
		Type:    1,
	}

	Users = append(Users, user)
	msg := &logic.Message{
		Content: nickname + "加入时间：" + time.Now().Format("yyyy/MM/dd HH:mm:ss"),
		Type:    2,
	}

	for _, user := range Users {
		user.MessageChannel <- msg
	}
	// var v interface{}
	// err = wsjson.Read(ctx, conn, &v)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	//conn.Close(websocket.StatusNormalClosure, "")

}
