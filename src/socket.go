package main

import (
    "fmt"
    "container/list"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    //"fmt"
    //"time"
)

const LongPolling int = 0

type Socket struct {
    Port int
    Token string
    Transport int
    Router *gin.Engine
    Clients chan chan string
    Connections map[string] interface{}
    Channels *list.List
    Template string
    Events *list.List
}

type Json map[string] interface{}

func (socket *Socket) Initialize() Socket {

    // Route
    gin.SetMode(gin.DebugMode)
    socket.Router = gin.Default()

    // Client channel
    socket.Clients = make(chan chan string, 1)

    // Initialize empty linked list
    socket.Channels = list.New()
    socket.Clients  = make(chan chan string, 1)

    // Socket template
    socket.Router.LoadHTMLGlob(socket.Template)

    // Events
    socket.Connections = make(map[string] interface{})

    return *socket
}

func (socket Socket) Debug(message string) {
    fmt.Println(message)
}

func (socket Socket) Emit(event string, data Json) {
    buffer, err := json.Marshal(data)

    if err != nil {
        panic(err)
    }
    fmt.Println(string(buffer[:]))
    //message <- string(buffer[:])
}

func (s Socket) Broadcast(event string, a interface {}) {
}

func (socket Socket) Wait(callback func()) {
    go func() {
        for {
            callback()
        }
    }()
}

func (socket Socket) On(event string, callback func(socket Socket)) {
    switch event {
    case "connection":
        socket.Wait(func() {
            select {
                case client := <- socket.Clients:
                    socket.Connections[random()] = client
                    callback(socket)
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
        Polling {
            Clients: socket.Clients,
            Context: context,
        }.Handle()
    })
    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}