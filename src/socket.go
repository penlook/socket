package socket

import (
    "fmt"
    "encoding/json"
)

const Polling int = 0

type Socket struct {
    Port int
    Transport int
}

type Json map[string] interface{}

func (s Socket) Handle() {
    fmt.Println("Handle")
}

func (s Socket) Emit(event string, data Json) {
    fmt.Println("Send event : " + event)
    buffer,err := json.Marshal(data)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(buffer[:]))
}

func (s Socket) Broadcast(event string, a interface {}) {
}

func (s Socket) On(event string, callback func()) {
}
