package main

import (
	"backend/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ws", controller.ChatWebSocketInit)
	r.Run(":2022")
}
