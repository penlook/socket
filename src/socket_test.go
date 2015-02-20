package main

import (
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

	fmt.Println("Initialize")
	socket.Initialize()

	socket.Static("/static", "./asset")
	socket.Router.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", Json {})
	})

	socket.On("connection", func(client Client) {

		client.Emit("abc", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		client.Emit("abc", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		client.Emit("abc", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		client.On("test", func(client Client) {
			client.Emit("test", Json {
				"abce" : "Hello",
				"abcf" : "Yes",
			})
		})
	})

	socket.Listen()
}


