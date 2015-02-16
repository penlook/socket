package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"time"
	//"fmt"
	//"io"
)

type Polling struct {
	Clients chan chan string
	Context *gin.Context
}

func (poll Polling) Handle() {
	message := make(chan string, 1)
	poll.Clients <- message

	go func() {
		time.Sleep(10 * time.Second)
		// Processing
		message <- "ABC"
	}()

	select {
		case <-time.After(time.Second * 20):
			poll.Context.JSON(200 , Json {
				"data" : "Timeout",
			})
		case msg := <-message:
			poll.Context.JSON(200 , Json {
				"data" : msg,
			})
	}
}