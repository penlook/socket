package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    //"fmt"
    //"time"
)

const LongPolling int = 0

type Json map[string] interface {}

type Context struct {
    Context *gin.Context
    Channel chan *gin.Context
    Output chan Json
    Handshake string
    Polling bool
}

type Socket struct {
    Port int
    Token string
    Transport int
 	Event map[string] interface {}
    Clients map[string] Client
    Context chan Context
    Router *gin.Engine
    Template string
}

func (socket *Socket) Initialize() Socket {

    // Route
    gin.SetMode(gin.DebugMode)
    socket.Router = gin.Default()

    // Context
    socket.Context = make(chan Context, 10)

    socket.Event = make(map[string] interface {})

    // Clients
    socket.Clients = make(map[string] Client)

    // Socket template
    socket.Router.LoadHTMLGlob(socket.Template)

    return *socket
}

func (socket Socket) Debug(message string) {
    fmt.Println(message)
}

func (socket Socket) ParseContext(context Context, callback func(client Client)) Client {

	var client Client
	if ! context.Polling {
		fmt.Println("No Handshake")
		client = socket.Clients[context.Handshake]
	    client.Emit("connection", Json {
	        "handshake" : context.Handshake,
	    })
	    fmt.Println("Callback")
	    defer callback(client)
	    return client
	}

	fmt.Println("Yes Handshake")

	// Update new context for client indentified by handshake id
	client = socket.Clients[context.Handshake]
	client.Context = context.Context

	if data := <- client.Output ; data != nil {
        client.Output <- data
    }

    return client
}

func (socket Socket) On(event string, callback func(client Client)) {
	socket.Event[event] = callback
}

func (socket Socket) Static(route string, directory string) Socket {
    socket.Router.Static(route, directory)
    return socket
}

// Check polling request per connection
func (socket Socket) LoopEvent(context Context) {
	go func(callback) {
		for {
			select {
				case context := <- context.Channel:
					socket.ParseContext(context, callback)
			}
		}
	}(socket.Event["connection"])
}

func (socket Socket) GetConnection(context *gin.Context) Context {

	handshake := random()
	output := make(chan Json, 10)

	client := Client {
       	Context: context,
       	Output: output,
    }

    socket.Clients[handshake] = client

	return Context {
		Context   : context,
		Output    : output,
		Handshake : handshake,
		Polling   : false,
	}
}

func (socket Socket) GetPolling(context *gin.Context) Context {
	fmt.Println("--> Get Polling")

	handshake := context.Params.ByName("handshake")

	client := socket.Clients[handshake]
	client.Context = context

	return Context {
		Context : context,
		Output : client.Output,
		Handshake: handshake,
		Polling: true,
	}
}

func (socket Socket) Response(context Context) {
	select {
		case data := <- context.Output:
			context.Context.JSON(200, data)
			fmt.Println("--> End Response")
	}
}

func (socket Socket) Listen() Socket {

    socket.Router.GET("/polling", func(_context *gin.Context) {
    	context := socket.GetConnection(_context)
    	socket.LoopEvent(context)
    	socket.Response(context)
    })

    socket.Router.GET("/polling/:handshake", func(_context *gin.Context) {
    	context := socket.GetPolling(_context)
    	context.Channel <- _context
    	socket.Response(context)
    })

    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}