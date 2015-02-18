package main

import (
    "fmt"
    "container/list"
    //"encoding/json"
    "github.com/gin-gonic/gin"
    //"github.com/tommy351/gin-sessions"
    "strconv"
    "net/http"
    //"fmt"
    //"time"
)

const LongPolling int = 0

type Json map[string] interface{}

type Socket struct {
    Port int
    Token string
    Transport int
    Router *gin.Engine
    Context chan *gin.Context
    Clients map[string] Client
    Connections map[string] interface {}
    Channels *list.List
    Output chan Json
    Template string
    Events *list.List
}

func (socket *Socket) Initialize() Socket {

    // Route
    gin.SetMode(gin.DebugMode)
    socket.Router = gin.Default()

    // Session
    /*store := sessions.NewCookieStore([]byte("secret123"))
    socket.Router.Use(sessions.Middleware("socket_session", store))*/

    // Context channel
    socket.Context = make(chan *gin.Context, 1)

    // Clients
    socket.Clients = make(map[string] Client, 10000)

    // Initialize empty linked list
    socket.Channels = list.New()

    // Socket template
    socket.Router.LoadHTMLGlob(socket.Template)

    // Events
    socket.Connections = make(map[string] interface{})

    // Signal
    socket.Output = make(chan Json, 100)

    return *socket
}

func (socket Socket) Debug(message string) {
    fmt.Println(message)
}

func (socket Socket) Wait(callback func()) {
    go func() {
        for {
            callback()
        }
    }()
}

func (socket Socket) On(event string, callback func(client Client)) {
    switch event {
    case "connection":
        socket.Wait(func() {
            select {
                case context := <- socket.Context:
                    context.Request.ParseForm()
                    handshake := context.Request.Form.Get("handshake")

                    var client Client

                    if handshake == "" {
                        handshake := random()
                        client = Client {
                            Context: context,
                            Output : socket.Output,
                            Handshake: handshake,
                        }
                        socket.Clients[handshake] = client
                        go func() {
                            client.Emit("connection", Json {
                                "handshake" : handshake,
                            })
                        }()
                        return
                    }

                    client = socket.Clients[handshake]

                    // Update new context
                    client.Context = context

                    callback(client)
            }
        })
    }
}

func (socket Socket) Static(route string, directory string) Socket {
    socket.Router.Static(route, directory)
    return socket
}

func (socket Socket) Listen() Socket {
    socket.Router.GET("/polling", func(context *gin.Context) {
        socket.Context <- context
        select {
        case data := <- socket.Output:
            context.JSON(200, data)
        }
    })
    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}