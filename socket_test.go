/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Loi Nguyen       <loint@penlook.com>
 */

package socket

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"io"
	"encoding/json"
	"bytes"
	"fmt"
)

var socket_socket = Socket {
	Port : 3000,
	Interval: 60,
}

func TestSocketInitialize(t *testing.T) {

	assert := assert.New(t)
	socket_socket.Initialize()
	assert.NotNil(socket_socket.Router)

	assert.NotNil(socket_socket.Context)
	assert.Equal(0, len(socket_socket.Context))

	assert.NotNil(socket_socket.Event)
	assert.Equal(0, len(socket_socket.Event))

	assert.NotNil(socket_socket.Clients)
	assert.Equal(0, len(socket_socket.Clients))

}

func TestSocketGetConnection(t *testing.T) {

	assert := assert.New(t)

	// Mockup HTTP Request
	request, _ := http.NewRequest("GET", "/polling_test_1", nil)

	// Create request recorder
	writer := httptest.NewRecorder()

	var context Context

	// Register handler for mock request
	socket_socket.Router.GET("/polling_test_1", func(context_ *gin.Context) {
		context = socket_socket.GetConnection(context_)
	})

	// Start request
	socket_socket.Router.ServeHTTP(writer, request)

	// Assert result
	assert.Equal(true, (len(context.Handshake) == 20))
	assert.Equal(false, context.Polling)
	assert.NotNil(socket_socket.Clients[context.Handshake])
}

func TestSocketGetPolling(t *testing.T) {

	assert := assert.New(t)

 	var handshake string

	for handshake_, _ := range socket_socket.Clients {
		handshake = handshake_
	}

	// Mockup HTTP Request

	request, _ := http.NewRequest("GET", "/polling_test_2/" + handshake, nil)
	writer := httptest.NewRecorder()

	var context Context

	// Register handler for mock request
	socket_socket.Router.GET("/polling_test_2/:handshake", func(context_ *gin.Context) {
		context = socket_socket.GetPolling(context_)
	})

	// Start request
	socket_socket.Router.ServeHTTP(writer, request)

	assert.Equal(true, (len(context.Handshake) == 20))
	assert.Equal(true, context.Polling)
	assert.Equal(socket_socket.Clients[handshake].Context, context.Context)
}

func TestSocketInitClientEvent(t *testing.T) {

	assert := assert.New(t)

 	var handshake string

	for handshake_, _ := range socket_socket.Clients {
		handshake = handshake_
	}

	request, _ := http.NewRequest("GET", "/polling_test_3/" + handshake, nil)
	writer := httptest.NewRecorder()

	var context Context

	// Register handler for mock request
	socket_socket.Router.GET("/polling_test_3/:handshake", func(context_ *gin.Context) {
		context = socket_socket.GetPolling(context_)
	})

	// Start request
	socket_socket.Router.ServeHTTP(writer, request)

	client := socket_socket.Clients[handshake]
	assert.Equal(false, client.HandshakeFlag)

	socket_socket.On("connection", func(client_ Client) {
		client_.On("event1", func(data Json) {})
		client_.On("event2", func(data Json) {})
		client_.On("event3", func(data Json) {})
		client_.On("event4", func(data Json) {})
	})

	socket_socket.InitClientEvent(context)

	client = socket_socket.Clients[handshake]
	assert.Equal(true, client.HandshakeFlag)

	// Fail
	//assert.Equal(4, client.MaxEvent)
}

// Test register event
func TestSocketOn(t *testing.T) {

	assert := assert.New(t)
	assert.Equal("test", "test")

	callback := func(client Client) {}
	socket_socket.On("connection", callback)
	socket_socket.On("disconnect", callback)

	assert.Equal(callback, socket_socket.Event["connection"])
	assert.Equal(callback, socket_socket.Event["disconnect"])
}

// Create mockup HTTP Request
func makeRequest(method, url string, data Json) *httptest.ResponseRecorder {
	query, _ := json.Marshal(data)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(query))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	socket_socket.Router.ServeHTTP(writer, request)
	return writer
}

// Convert string to JSON type
func toJson(data io.Reader) Json {
	decoder := json.NewDecoder(data)
	var json Json
	decoder.Decode(&json)
	return json
}

// Integration Test
//
// Step 1: Initialize new polling connection
// Step 2: Setup event in server
// Step 3: Emit new event to server
// Step 4: Client recieve event from server
func TestSocketClientServer(t *testing.T) {
	assert := assert.New(t)

	// Step 1
	// Initialize polling connection
	socket_socket.Router.GET("/polling_test_socket_1", socket_socket.ServePooling())
	response := makeRequest("GET", "/polling_test_socket_1", Json {})
	assert.NotNil(response)

	data := toJson(response.Body)
	assert.NotNil(data)

	handshake := data["data"].(map[string] interface {})["handshake"]
	assert.Equal(true, len(handshake) == 20)

	// Step 2
	// Install event in server
	socket_socket.On("connection", func(client Client) {
		client.On("")
	})

	// Using exist handshake to emit new data
	socket_socket.Router.GET("/polling_test_socket_2/:handshake", socket_socket.ServePostHandshake())
	response := makeRequest("GET", "/polling_test_socket_1", Json {
		"event" : ""
	})

	/*
	socket := Socket {
		Port: 3000,
		Token: "acbz@3345123124567",
		Interval: 60,
	}

	socket.Initialize()
	socket.Static("/static", "./example")

	socket.Template("example")
	socket.Router.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", Json {})
	})

	socket.On("connection", func(client Client) {
		client.On("init", func(data Json) {

			client.Broadcast("test", Json {
				"eventdata" : "Broadcast",
			})

			client.On("test2", func(data Json) {
				client.Emit("test2", Json {
					"key" : "Package 2 from server",
				})
			})
		})
		client.Emit("test_on", Json {
			"key" : "abcxyz",
		})
		client.On("test", func(data Json) {
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
	*/
}


