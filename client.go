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
)

// Client structure
type Client struct {
	Socket Socket
	Context *gin.Context
	Channel chan Context
	Output chan Json
	Handshake string
	HandshakeFlag bool
	Event *list.List
	MaxEvent int
}

// Listen event on client
//
// client.On("event", func(client Client) {
// 		// TODO
// })
func (client Client) On(event string, callback func(data Json)) {
	client.MaxEvent = client.MaxEvent + 1
	client.Event.PushBack( Event {
		Id : client.MaxEvent,
		Name : event,
		Callback : callback,
	})
}

// Push event to client
//
// client.Emit("event", Json {
// 		"key1" : "value1",
// 		"key2" : "value2",
// })
func (client Client) Emit(event string, data Json) {
	client.Output <- Json {
    	"event": event,
    	"data" : data,
    }
}

// Broadcast event to otherwise client
//
// client.Broadcast("event", Json {
// 		"key1" : "value1",
// 		"key2" : "value2",
// })
func (client Client) Broadcast(event string, data Json) {
	for handshake, client_ := range client.Socket.Clients {

		// Parallel Broadcasting
		go func(handshake string, client_ Client, event string, data Json) {
			if handshake != client.Handshake {
				client_.Output <- Json {
					"event": event,
					"data" : data,
				}
			}
		} (handshake, client_, event, data)

	}
}

// Broadcast event to all client
//
// client.BroadcastAll("event", Json {
// 		"key1" : "value1",
// 		"key2" : "value2",
// })
func (client Client) BroadcastAll(event string, data Json) {
	for handshake, client_ := range client.Socket.Clients {

		// Parallel Broadcasting
		go func(handshake string, client_ Client, event string, data Json) {
			client_.Output <- Json {
					"event": event,
					"data" : data,
			}
		} (handshake, client_, event, data)

	}
}
