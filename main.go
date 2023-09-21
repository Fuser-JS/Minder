package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Item struct {
	ID uint
	item string
}

// env loading
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


// Database Initialization
func dbInit() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
}


func main() {
	dbInit()
	r:=gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
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
