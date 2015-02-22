package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oleiade/lane"
	"fmt"
	//"container/list"
)

type Event struct {
	Name string
	Callback func(client Client)
}

type Client struct {
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	Event *lane.Stack
	MaxNode int
}

func (client Client) On(event string, callback func(client Client)) {
	fmt.Println("Client On " + event)
	client.MaxNode = client.MaxNode + 1

	client.Event.Push( Node{
		Id : client.MaxNode,
		Event : event,
		Callback : callback,
	})

	fmt.Println("Recursive")
	callback(client)

	fmt.Println("Back tracking")
	/*node := client.Event.Pop()
	fmt.Println(node)*/
}

func (client Client) Emit(event string, data Json) {
	/*fmt.Println(client.Event)
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }*/
}
