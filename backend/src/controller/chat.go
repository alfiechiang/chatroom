package controller

import (
	"backend/src/logic"
	"context"

	// "fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
	// "nhooyr.io/websocket/wsjson"
)


//	users: make(map[string]*User),

var Users map[string]*logic.User = make(map[string]*logic.User)
var brodcast = logic.NewBrodcast()


func ChatWebSocketInit(c *gin.Context) {

	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*30)
	defer cancel()

	nickname := c.Request.FormValue("nickname")

	user := logic.NewUser(nickname, conn, ctx)
	go user.SendMessage(c)

	go brodcast.Start()

	user.MessageChannel <- logic.WelcomeMessage(nickname)

	brodcast.Users = append(brodcast.Users, user)
	msg := logic.InviteMessage(nickname)
	brodcast.MessageChannel <- msg

	brodcast.ReceiveMessage(user)

	//conn.Close(websocket.StatusNormalClosure, "")

}
