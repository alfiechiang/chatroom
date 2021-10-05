package controller



import (
	"context"
	"fmt"
	"log"

	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	MessageChannel chan string `json:"-"`
	Nickname string
}
var Users []User = make([]User,0)

func ChatWebSocketInit(c *gin.Context){
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			log.Println(err)
			return
		}
		nickname := c.Request.FormValue("nickname")
		defer conn.Close(websocket.StatusInternalError, "Internal Error")

		ctx, cancel := context.WithTimeout(c, time.Second*30)
		defer cancel()
		fmt.Println(nickname)
		wsjson.Write(ctx, conn, nickname+"您好，欢迎加入聊天室！")

		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Close(websocket.StatusNormalClosure, "")

}