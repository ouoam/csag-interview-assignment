package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type mydata struct {
	Age      int
	Name     string
	Lastname string
}

type myTime struct {
	ISO       time.Time
	Stamp     int64
	StampNano int64
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		t := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"Me": &mydata{
				Age:      19,
				Name:     "Phumphathai",
				Lastname: "Chansriwong",
			},

			"Time": &myTime{
				ISO:       t,
				Stamp:     t.Unix(),
				StampNano: t.UnixNano(),
			},

			"Request.Header": c.Request.Header,
		})
	})

	router.Run(":" + port)
}
