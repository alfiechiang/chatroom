package main

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

func main() {

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "Internal Error")

		ctx, cancel := context.WithTimeout(c, time.Second*30)
		defer cancel()
		fmt.Println("Hello World")
		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Close(websocket.StatusNormalClosure, "")
	})

	r.Run(":2022")

	

}
