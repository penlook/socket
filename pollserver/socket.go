package main

import (
	. "github.com/penlook/socket"
)

func main() {

	socket := Socket {
		Port: 1234,
		Token: "acbz@3345123124567",
		Interval: 10,
	}

	socket.Initialize()

	// Mapping resources
	socket.Static("/client", "./../client")

	socket.Listen()
}