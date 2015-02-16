package main

import (
    "fmt"
    "container/list"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    //"time"
)

const LongPolling int = 0

type Socket struct {
    Port int
    Token string
    Transport int
    Router *gin.Engine
    Clients chan chan string
    Channels *list.List
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

    return *socket
}

func (s Socket) Emit(event string, data Json) {
    fmt.Println("Send event : " + event)
    buffer, err := json.Marshal(data)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(buffer[:]))
}

func (s Socket) Broadcast(event string, a interface {}) {
}

func (socket Socket) On(event string, callback func()) Socket {
    return socket
}

func (socket Socket) Static(route string, directory string) Socket {
    socket.Router.Static(route, directory)
    return socket
}

func (socket Socket) ClientHandler() {
    go func() {
        for {
            select {
            case client := <- socket.Clients:
                socket.Channels.PushBack(client)
                fmt.Printf("New client: %d\n", socket.Channels.Len())
            }
        }
    }()
}

func (socket Socket) Listen() Socket {
    socket.ClientHandler()
    socket.Router.GET("/polling", func(context *gin.Context) {
        Polling {
            Clients: socket.Clients,
            Context: context,
        }.Handle()
    })
    http.ListenAndServe(":" + strconv.Itoa(socket.Port), socket.Router)
    return socket
}