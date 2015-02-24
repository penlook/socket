package main

import (
	. "github.com/penlook/socket"
)

func main() {
	socket := Socket {
	    Port: 3000,
	    Token: "acbz@3345123124567",
	    Transport: LongPolling,
	    Template: "asset/*",
	}

	socket.Initialize()
	socket.Static("/static", "./asset")

	socket.On("connection", func(client Client) {
	    client.On("init", func(data Json) {
	    })
	})

	socket.Listen()
}

