package socket

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

// Event structure
type Event struct {
	Id int
	Name string
	Callback func(data Json)
}

// Client structure
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

// Listen event on client
//
// client.On("event", func(client Client) {
// 		// TODO
// })
func (client Client) On(event string, callback func(data Json)) {
	client.MaxEvent = client.MaxEvent + 1
	client.Event.PushBack( Event {
		Id : client.MaxEvent,
		Name : event,
		Callback : callback,
	})
}

// Push event to client
//
// client.Emit("event", Json {
// 		"key1" : "value1",
// 		"key2" : "value2",
// })
func (client Client) Emit(event string, data Json) {
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}

// Broadcast event to otherwise client
//
// client.Broadcast("event", Json {
// 		"key1" : "value1",
// 		"key2" : "value2",
// })
func (client Client) Broadcast(event string, data Json) {
	for handshake, client_ := range client.Socket.Clients {

		// Parallel Broadcasting
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
