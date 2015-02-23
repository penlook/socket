package main

import (
    "fmt"
    "container/list"
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    "encoding/json"
    //"fmt"
    //"time"
)

const LongPolling int = 0

type Json map[string] interface {}

type Context struct {
    Context *gin.Context
    Channel chan Context
    Output chan Json
    Handshake string
    Polling bool
}

type Socket struct {
    Port int
    Token string
    Transport int
 	Event map[string] func(client Client)
    Clients map[string] Client
    Context chan Context
    Output chan Json
    Router *gin.Engine
    Template string
}

func (socket *Socket) Initialize() Socket {

    // Route
    gin.SetMode(gin.DebugMode)
    socket.Router = gin.Default()

    // Context
    socket.Context = make(chan Context, 10)

 	// Event
    socket.Event = make(map[string] func(client Client))

    // Clients
    socket.Clients = make(map[string] Client)

    // Output
    socket.Output = make(chan Json, 100)

    // Socket template
    socket.Router.LoadHTMLGlob(socket.Template)

    return *socket
}

func (socket Socket) Debug(message string) {
    fmt.Println(message)
}

func (socket Socket) UpdateContext(context Context) Client {

	client := socket.Clients[context.Handshake]
    client.Context = context.Context

	if ! context.Polling {
        socket.Emit("connection", Json {
	        "handshake" : context.Handshake,
	    })
	}

    return client
}

func (socket Socket) On(event string, callback func(client Client)) {
	socket.Event[event] = callback
}

func (socket Socket) Emit(event string, data Json) {
    fmt.Println("Socket emit handshake")
	socket.Output <- Json {
    	"event": event,
    	"data" : data,
    }
    fmt.Println("Done emit")
}

func (socket Socket) Static(route string, directory string) Socket {
    socket.Router.Static(route, directory)
    return socket
}

// Check polling request per connection
func (socket Socket) LoopSocketEvent(context Context) {
    go func(context Context) {
		for {
			select {
				case context := <- context.Channel:
					socket.UpdateContext(context)
			}
		}
	} (context)
}

func (socket Socket) InitClientEvent(context Context) {
    client := socket.Clients[context.Handshake]

    if ! client.HandshakeFlag {
        fmt.Println("Event initialize")
        client.HandshakeFlag = true
        socket.Clients[context.Handshake] = client
        callback := socket.Event["connection"]
        callback(client)
    }

    fmt.Println(" -------- TOTAL EVENT IS -------- ")
    var node Node
    for cursor := client.Event.Front(); cursor != nil; cursor = cursor.Next() {
        node = cursor.Value.(Node)
        fmt.Println(node.Event)
        fmt.Println("\n")
    }
}

func (socket Socket) SubmitClientEvent(context Context) {

    decoder := json.NewDecoder(context.Context.Request.Body)
    var pkg Json
    decoder.Decode(&pkg)
    event_name := pkg["event"]
    client := socket.Clients[context.Handshake]

    var node Node
    for cursor := client.Event.Front(); cursor != nil; cursor = cursor.Next() {
        node = cursor.Value.(Node)
        if node.Event == event_name {
            node.Callback(pkg)
        }
    }

    fmt.Println(" -------- TOTAL EVENT IS -------- ")
    for cursor := client.Event.Front(); cursor != nil; cursor = cursor.Next() {
        node = cursor.Value.(Node)
        fmt.Println(node.Event)
        fmt.Println("\n")
    }

    context.Context.JSON(200, Json {
         "status" : "OK",
    })
}

func (socket Socket) GetConnection(context *gin.Context) Context {

	handshake := random()
	output    := make(chan Json, 10)
	channel   := make(chan Context, 10)
	event     := list.New()

	client := Client {
       	Context: context,
       	Output : output,
       	Channel: channel,
        Event: event,
        HandshakeFlag: false,
        MaxNode: 0,
    }

    socket.Clients[handshake] = client

	return Context {
		Context   : client.Context,
		Output    : client.Output,
		Channel   : client.Channel,
		Handshake : handshake,
		Polling   : false,
	}
}

func (socket Socket) GetPolling(context *gin.Context) Context {
    fmt.Println("Get Polling")
	handshake := context.Params.ByName("handshake")

	client := socket.Clients[handshake]
	client.Context = context

    fmt.Println("Return context")

	return Context {
		Context : context,
		Channel : client.Channel,
		Output : client.Output,
		Handshake: handshake,
		Polling: true,
	}
}

func (socket Socket) Response(context Context) {
	select {
        case data := <- socket.Output:
            context.Context.JSON(200, data)
        case data := <- context.Output:
            context.Context.JSON(200, data)
	}
}

func (socket Socket) Listen() Socket {

    socket.Router.GET("/polling", func(_context *gin.Context) {
    	context := socket.GetConnection(_context)
    	socket.LoopSocketEvent(context)
    	context.Channel <- context
    	socket.Response(context)
    })

    socket.Router.GET("/polling/:handshake", func(_context *gin.Context) {
    	context := socket.GetPolling(_context)
        socket.InitClientEvent(context)
    	context.Channel <- context
    	socket.Response(context)
    })

    socket.Router.POST("/polling/:handshake", func(_context *gin.Context) {
        context := socket.GetPolling(_context)
        socket.SubmitClientEvent(context)
        context.Channel <- context
    })

    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}
