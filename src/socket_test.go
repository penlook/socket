package main

import	(
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSocket(t *testing.T) {

	assert := assert.New(t)
	assert.Equal("Test", "Test")

	socket := Socket {
		Port: 3000,
		Token: "acbz@3345123124567",
		Transport: LongPolling,
		Template: "asset/*",
	}

	socket.Initialize()
	socket.Static("/static", "./asset")

	socket.On("connection", func(socket Socket) {
		socket.Debug("Client connected !")
		socket.Emit("init", Json {
			"user": "loint",
			"token": "abcdd123sdc",
		})
	})

	socket.Listen()
}


