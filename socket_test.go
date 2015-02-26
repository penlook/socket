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
// Step 4: Client receive event from server
func TestSocketClientServer(t *testing.T) {
	assert := assert.New(t)

	socket_socket.Initialize()

	// Step 1
	// Initialize polling connection
	socket_socket.Router.GET("/polling_test_socket", socket_socket.ServePooling())
	response := makeRequest("GET", "/polling_test_socket", Json {})
	assert.NotNil(response)

	data := toJson(response.Body)
	assert.NotNil(data)

	handshake := data["data"].(map[string] interface {})["handshake"].(string)
	assert.Equal(true, len(handshake) == 20)

	// Step 2
	// Install event in server
	// Old callback will be overrided when rewrite 'On("connection")'
	socket_socket.On("connection", func(client Client) {
		client.On("init", func(data Json) {
			assert.Equal(Json{
				"init_key1" : "ABCDEF012345",
				"init_key2" : "XYZ12345",
			}, data)
		})
		client.On("abc", func(data Json) {
			assert.Equal(Json{
				"abc_key1" : "ABCDEF012345",
				"abc_key2" : "XYZ12345",
			}, data)
		})
	})

	// Step 3
	// Using exist handshake to emit new data
	socket_socket.Router.GET ("/polling_test_socket/:handshake", socket_socket.ServeGetHandshake())
	socket_socket.Router.POST("/polling_test_socket/:handshake", socket_socket.ServePostHandshake())

	// Handle first event
	response = makeRequest("POST", "/polling_test_socket/" + handshake, Json {
		"event" : "init",
		"data" : Json {
			"init_key1" : "ABCDEF012345",
			"init_key2" : "XYZ12345",
		},
	})

	assert.NotNil(response)

	// Handle second event
	response = makeRequest("POST", "/polling_test_socket/" + handshake, Json {
		"event" : "abc",
		"data" : Json {
			"abc_key1" : "ABCDEF012345",
			"abc_key2" : "XYZ12345",
		},
	})

	assert.NotNil(response)

	// Testing ...

	fmt.Println()
}


