package main

import (
    "fmt"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "strconv"
)

const Polling int = 0

type Socket struct {
    Port int
    Token string
    Transport int
    Router *gin.Engine
}

type Json map[string] interface{}

func (socket *Socket) Initialize() Socket {

    // Route
    gin.SetMode(gin.DebugMode)
    socket.Router = gin.Default()

    // Anything else ?

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

func (socket Socket) Listen() Socket {
    socket.Router.Run(":" + strconv.Itoa(socket.Port))
    return socket
}