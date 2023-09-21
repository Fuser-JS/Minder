package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	
	r.GET("/", func(c *gin.Context) {
		c.File("video/rick.mp4")
	})
	
	r.Run()
}
