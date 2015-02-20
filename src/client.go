package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	//"container/list"
)

type Event struct {
	Name string
	Callback func(client Client)
	//Events *List
}

type Client struct {
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	//Events *List
}

func (client Client) On(event string, callback func(client Client)) {
	//a := List.New()
	//fmt.Println(a)
	fmt.Println("Run event " + event)
	callback(client)
}

func (client Client) Emit(event string, data Json) {
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}