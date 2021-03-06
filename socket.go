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
	"container/list"
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
	"time"
	//"fmt"
)

// Long polling - Implemented
const LongPolling int = 0

// Web socket - Not yet implemented
const WebSocket int = 1

// Server sent - Not yet implemented
const Serversent int = 2

// Request context structure
type Context struct {
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	Polling bool
}

// Socket structure
type Socket struct {
	Port int
	Token string
	Transport int
	Interval time.Duration
	Event map[string] func(client Client)
	Clients map[string] *Client
	Context chan Context
	Router *gin.Engine
}

// Initialize for socket
func (socket *Socket) Initialize() Socket {

	// Route using Gin Framework
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	socket.Router = gin.Default()

	// Context
	socket.Context = make(chan Context)

	// Event
	socket.Event = make(map[string] func(client Client))

	// Clients
	socket.Clients = make(map[string] *Client)

	return *socket
}

// Update new context for new or exist client per handshake
func (socket Socket) UpdateContext(context Context) Client {

	client := socket.Clients[context.Handshake]
	client.Context = context.Context

	if ! context.Polling {
		context.Context.JSON(200, Json {
			"event" : "connection",
			"data"  : Json {
				"handshake" : context.Handshake,
			},
		})
	}

	return *client
}

// Socket listen client event
func (socket Socket) On(event string, callback func(client Client)) {
	socket.Event[event] = callback
}

// Using template for HTML Output
func (socket Socket) Template(template_directory string) {
	socket.Router.LoadHTMLGlob(template_directory + "/*")
}

// Static resources
func (socket Socket) Static(route string, directory string) Socket {
	socket.Router.Static(route, directory)
	return socket
}

// Scan client events in the first handshake
func (socket Socket) InitClientEvent(context Context) {

	client  := socket.Clients[context.Handshake]

	if ! client.HandshakeFlag {
		client.HandshakeFlag = true
		socket.Clients[context.Handshake] = client
		callback := socket.Event["connection"]
		callback(*client)
	}
}

// Submit emit package from client
func (socket Socket) SubmitClientEvent(context Context) {

	// Decode request data to Json type
	decoder := json.NewDecoder(context.Context.Request.Body)
	var pkg Json
	decoder.Decode(&pkg)
	event_name := pkg["event"]

	// Temporary type casting
	event_data_raw := pkg["data"].(map[string] interface{})

	// Convert from map[string] interface {} to Json type
	event_data := Json {}
	for key, value := range event_data_raw {
		event_data[key] = value
	}

	// Get current client based on handshake id
	client := socket.Clients[context.Handshake]

	var event Event

	// Scan all event and callback correspoding event
	for cursor := client.Event.Front(); cursor != nil; cursor = cursor.Next() {
		event = cursor.Value.(Event)
		if event.Name == event_name {
			event.Callback(event_data)
		}
	}

	context.Context.JSON(200, Json {
		"status" : "OK",
	})
}

// Get fresh connection
func (socket Socket) GetConnection(context *gin.Context) Context {

	handshake := random()
	output    := make(chan Json, 10)
	channel   := make(chan Context, 10)
	event     := list.New()

	client := Client {
		Socket: socket,
		Output : output,
		Context: context,
		Channel: channel,
		Event: event,
		Handshake: handshake,
		HandshakeFlag: false,
		MaxEvent: 0,
	}

	socket.Clients[handshake] = &client

	return Context {
		Context   : client.Context,
		Output    : client.Output,
		Channel   : client.Channel,
		Handshake : handshake,
		Polling   : false,
	}
}

// Get polling connection
func (socket Socket) GetPolling(context *gin.Context) Context {
	handshake := context.Params.ByName("handshake")
	client := socket.Clients[handshake]
	client.Context = context

	return Context {
		Context : context,
		Channel : client.Channel,
		Output : client.Output,
		Handshake: handshake,
		Polling: true,
	}
}

// Waiting for long-polling response
func (socket Socket) Response(context Context) {
	timeout := make(chan bool, 1)

	// Timeout monitoring
	go func() {
		time.Sleep(socket.Interval * time.Second)
		timeout <- true
	}()

	select {
		case data := <- context.Output:
			context.Context.JSON(200, data)
		case <- timeout:
			context.Context.JSON(200, Json {
				"event" : "heartbeat",
				"status" : "good",
			})
	}
}

func (socket Socket) ServePooling() gin.HandlerFunc {
	return func(_context *gin.Context) {
		context := socket.GetConnection(_context)
		socket.UpdateContext(context)
	}
}

func (socket Socket) ServeGetHandshake() gin.HandlerFunc {
	return func(_context *gin.Context) {
		context := socket.GetPolling(_context)
		socket.InitClientEvent(context)
		context.Channel <- context
		socket.Response(context)
	}
}

func (socket Socket) ServePostHandshake() gin.HandlerFunc {
	return func(_context *gin.Context) {
		context := socket.GetPolling(_context)
		socket.SubmitClientEvent(context)
		context.Channel <- context
	}
}

func (socket Socket) SetAllowCrossDomain() {
	socket.Router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Next()
    })
}

// Listen HTTP Request
func (socket Socket) Listen() Socket {
	socket.SetAllowCrossDomain()

	socket.Router.GET ("/polling"           , socket.ServePooling())
	socket.Router.GET ("/polling/:handshake", socket.ServeGetHandshake())
	socket.Router.POST("/polling/:handshake", socket.ServePostHandshake())

	socket.Router.Run(":" + strconv.Itoa(socket.Port))
	return socket
}
