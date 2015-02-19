package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/session"
	//"time"
	"fmt"
)

type Client struct {
	Context *gin.Context
	Channel chan Json
	Handshake string
	Timeout int32
}

func (client Client) Emit(event string, data Json) {
	fmt.Println("Before EMIT " + event)
    client.Channel <- Json {
    	"event": event,
    	"data" : data,
    }
    fmt.Println("After EMIT " + event)
    fmt.Println("Returned")
}


