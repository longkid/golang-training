package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/longkid/golang-training/assignment/producer"
)

type Message struct {
	Message string `json:"message" binding:"required"`
}

func SendMessage(c *gin.Context) {
	// Get message from query param 'message'
	var json Message
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// TODO How to parameterize topic "golang-assignment-events"?
	err := producer.Publish("golang-assignment-events", json.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": "Publish message to Kafka topic fails"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Sending message succeeds"})
}
