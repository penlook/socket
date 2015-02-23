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

	fmt.Println("On connection")
	socket.On("connection", func(client Client) {
		fmt.Println("Emit abc")
		client.Emit("abc1", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		client.Emit("abc2", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		client.Emit("abc3", Json {
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",
		})

		fmt.Println("Client on")
		client.On("abc", func(data Json) {
			client.Emit("abc", Json {
				"abced": "1234",
			})
		})
		client.On("test", func(data Json) {
			fmt.Println("Enter test event")
			client.Emit("test", Json {
				"abce" : "Hello",
				"abcf" : "Yes",
			})

			client.On("test_2", func(data Json) {
				client.Emit("abc", Json {
					"test" : "1234",
				})
				client.On("test_3", func(data Json) {
					client.On("test_4", func(data Json) {
						client.Emit("abc", Json {
							"test" : "1234",
						})
					})
					client.Emit("abc", Json {
						"test" : "test",
					})
				})
				client.On("test_5", func(data Json) {
					client.Emit("abc", Json {
						"abc3" : "234556",
					})
					client.On("test7", func(data Json) {
						client.On("test8", func(data Json) {
							client.Emit("abc", Json {
								"abc" : "avc",
							})
						})
						client.On("test9", func(data Json) {
							client.Emit("abc", Json {
								"abc" : "avc",
							})
						})
					})
				})
				client.Emit("abc", Json {
					"abc" : "abc",
				})
			})
			client.On("test6", func(data Json) {
				client.Emit("abcer", Json {
					"abc" : "1245667",
				})
			})
		})
	})

	socket.Listen()
}


