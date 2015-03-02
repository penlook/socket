package main

import (
	daemon "github.com/penlook/daemon"
	. "github.com/penlook/socket"
)

func main() {
	daemon.Service {
		Name: "socket",
		Description: "Socket service",
		Process: Service,
		Port: 1234,
	}.Initialize()
}

func Service(server daemon.Service) {

	socket := Socket {
		Port: server.Port,
		Token: "acbz@3345123124567",
		Interval: 10,
	}

	socket.Initialize()

	// Mapping resources
	socket.Static("/client", "./../client")

	socket.Listen()
}