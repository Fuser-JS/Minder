package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// env loading
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}




func main() {
	r:=gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.1/24"})
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
