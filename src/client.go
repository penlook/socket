package main

import (
	"github.com/gin-gonic/gin"
)

type Client struct {
	Context *gin.Context
	Output chan Json
	Handshake string
}

func (client Client) On(event string, callback func(client Client)) {

}

func (client Client) Emit(event string, data Json) {
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}