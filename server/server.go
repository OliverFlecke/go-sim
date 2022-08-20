// package server

package main

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	stream := make(chan time.Time)
	go func() {
		defer close(stream)
		for {
			stream <- time.Now()
			time.Sleep(1 * time.Second)
		}
	}()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-stream; ok {
			c.SSEvent("message", msg)
			return true
		}

		return false
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/stream", StreamHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
