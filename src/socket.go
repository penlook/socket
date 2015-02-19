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

type Context struct {
    Context *gin.Context
    Channel chan Json
}

type Socket struct {
    Port int
    Token string
    Transport int
    Router *gin.Engine
    Context chan Context
    Clients map[string] Client
    Client chan Client
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

    // Request context
    socket.Context = make(chan Context, 1000)

    // Clients
    socket.Clients = make(map[string] Client, 1000)
    socket.Client = make(chan Client, 1000)

    // Initialize empty linked list
    socket.Channels = list.New()

    // Socket template
    socket.Router.LoadHTMLGlob(socket.Template)

    // Events
    socket.Connections = make(map[string] interface{})

    // Signal
    socket.Output = make(chan Json, 1000)

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
                    fmt.Println("Process context")
                    context.Context.Request.ParseForm()
                    handshake := context.Context.Request.Form.Get("handshake")

                    var client Client

                    if handshake == "" {
                        fmt.Println("No Handshake")
                        handshake := random()
                        fmt.Println("New Client " + handshake)

                        client = Client {
                            Context: context.Context,
                            Channel: context.Channel,
                            Handshake: handshake,
                        }

                        socket.Clients[handshake] = client
                        client.Emit("connection", Json {
                            "handshake" : handshake,
                        })
                        fmt.Println("Callback")
                        callback(client)
                        return
                    }

                    fmt.Println("YES Handshake")

                    // Update new context
                    client = socket.Clients[handshake]
                    client.Context = context.Context

                    fmt.Println(" --> TEST CHANNEL ")
                    if data := <- client.Channel ; data != nil {
                        socket.Output <- data
                    }

                    //client.Context.JSON(200, data)
                    //break
            }
        })
    }
}

func (socket Socket) Static(route string, directory string) Socket {
    socket.Router.Static(route, directory)
    return socket
}

func (socket Socket) Listen() Socket {

    socket.Router.GET("/polling", func(_context *gin.Context) {
        fmt.Println("----------- New Request ----------")
        context := Context {
            Context : _context,
            Channel : make(chan Json, 100),
        }
        fmt.Println("----------- Assign context ----------")
        socket.Context <- context

        fmt.Println("----------- Select channel ----------")

        select {
        case data := <- context.Channel:
            context.Context.JSON(200, data)
            fmt.Println(" ---------- Request End -------- ")
        case data := <- socket.Output:
            context.Context.JSON(200, data)
            fmt.Println(" ---------- Request End -------- ")
        }
    })

    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}