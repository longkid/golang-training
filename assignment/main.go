package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longkid/golang-training/assignment/controller"
)

func main() {
	r := gin.Default()

	r.POST("/send-message", controller.SendMessage)

	r.Run() // listen and serve on 0.0.0.0:8080
}
