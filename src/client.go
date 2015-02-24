package main

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

type Event struct {
	Name string
	Callback func(data Json)
}

type Client struct {
	Socket Socket
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	HandshakeFlag bool
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

func (client Client) Broadcast(event string, data Json) {
	for handshake, _ := range client.Socket.Clients {
		go func(handshake string, event string, data Json) {
			if handshake != client.Handshake {
				client_ := client.Socket.Clients[handshake]
				client_.Output <- Json {
					"event": event,
					"data" : data,
				}
			}
		} (handshake, event, data)
	}
}
