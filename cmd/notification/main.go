package main

import (
	"github.com/gin-gonic/gin"
	"github.com/talbx/mqtt-wire/internal/app/notification"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var input notification.DeviceData
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Fatal(err)
			return
		}
		service := notification.PushNotificationService{}
		var processor = notification.StatusProcessor{NotificationService: service}
		resp := processor.ProcessStatus(input)
		if resp != nil {
			c.JSON(http.StatusOK, resp)
			return
		}

		c.JSON(http.StatusOK, "No message sent, all okay")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
