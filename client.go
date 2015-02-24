package socket

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

type Event struct {
	Id int
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
	MaxEvent int
}

func (client Client) On(event string, callback func(data Json)) {
	client.MaxEvent = client.MaxEvent + 1
	client.Event.PushBack( Event {
		Id : client.MaxEvent,
		Name : event,
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
	for handshake, client_ := range client.Socket.Clients {
		go func(handshake string, client_ Client, event string, data Json) {
			if handshake != client.Handshake {
				client_.Output <- Json {
					"event": event,
					"data" : data,
				}
			}
		} (handshake, client_, event, data)
	}
}
