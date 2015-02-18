package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/session"
	//"time"
	//"fmt"
)

type Client struct {
	Context *gin.Context
	Output chan Json
	Handshake string
	Timeout int32
}

func (client Client) Emit(event string, data Json) {
    client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}


