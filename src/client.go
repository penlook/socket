package main

import (
	"container/list"
	"github.com/gin-gonic/gin"
	//"fmt"
)

type Event struct {
	Name string
	Callback func(data Json)
}

type Client struct {
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	Event *list.List
	MaxNode int
}

func (client Client) On(event string, callback func(data Json)) {
	client.MaxNode = client.MaxNode + 1
	client.Event.PushBack(Node {
		Id : client.MaxNode,
		Event : event,
		Callback : callback,
	})
}

func (client Client) Emit(event string, data Json) {
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}
