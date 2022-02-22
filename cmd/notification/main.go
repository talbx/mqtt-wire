package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	service "github.com/talbx/mqtt-wire/internal/app/notification"
	"github.com/talbx/mqtt-wire/internal/pkg/utils"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var input service.DataPoint
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Fatal(err)
			return
		}

		sprintf := fmt.Sprintf("Topic is %s, Battery status is %d", input.Message.Topic, input.Message.BatteryStatus)
		utils.LogStr(sprintf)
		c.JSON(http.StatusOK, input)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
