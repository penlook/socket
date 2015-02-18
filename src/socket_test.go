package main

import	(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"fmt"
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
	socket.Router.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", Json {})
	})

	socket.On("connection", func(client Client) {
		fmt.Println("Init")
		client.Emit("init", Json {
			"user": "loint",
			"token": "abcdd123sdc",
		})
		fmt.Println("Init2")
		client.Emit("init2", Json {
			"user": "l234234",
			"token": "abcd23432sdc",
		})
	})

	socket.Listen()
}


