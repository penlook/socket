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
	"container/list"
	"time"
)

var socket = Socket {
		Port : 3000,
		Interval: 60,
}

func createClient() Client {
	handshake := random()
	output    := make(chan Json, 10)
	channel   := make(chan Context, 10)
	event     := list.New()

	client := Client {
		Socket: socket,
		Context: nil,
		Output : output,
		Channel: channel,
		Event: event,
		Handshake: handshake,
		HandshakeFlag: false,
		MaxEvent: 0,
	}

	return client
}

func TestClientOn(t *testing.T) {

	assert := assert.New(t)
	client := createClient()

	assert.NotNil(client)

	client.On("test", func(data Json) {
	})

	client.On("abc", func(data Json) {
	})

	assert.Equal(2, client.MaxEvent)
	assert.Equal(2, client.Event.Len())
}

func TestClientEmit(t *testing.T) {

	assert := assert.New(t)
	client := createClient()

	assert.NotNil(client)

	client.Emit("test", Json {
		"key1" : "value1",
		"key2" : "value2",
	})

	client.Emit("abc", Json {
		"key1" : "value1",
		"key2" : "value2",
	})

	assert.Equal(Json {
		"event" : "test",
		"data"	: Json {
			"key1" : "value1",
			"key2" : "value2",
		},
	}, <- client.Output)

	assert.Equal(Json {
		"event" : "abc",
		"data"	: Json {
			"key1" : "value1",
			"key2" : "value2",
		},
	}, <- client.Output)

}

func TestClientBroadcast(t *testing.T) {

	assert := assert.New(t)
	client := createClient()

	assert.NotNil(client)
	socket.Initialize()
	assert.NotNil(socket.Clients)
	assert.Equal(0, len(socket.Clients))
	socket.Clients[client.Handshake] = client

	num := 10000

	for i := 0; i < num; i++ {
		client_ := createClient()
		socket.Clients[client_.Handshake] = client_
	}

	client.Socket = socket

	assert.Equal(num + 1, len(client.Socket.Clients))

	client.Broadcast("test", Json {
		"key1" : "value1",
		"key2" : "value2",
	})

	times := 0

	for handshake, client_ := range socket.Clients {
		if client.Handshake != handshake {
			go func(client_ Client) {
				times ++
				assert.Equal(Json {
					"event" : "test",
					"data"  : Json {
						"key1" : "value1",
						"key2" : "value2",
					},
				}, <- client_.Output)
			} (client_)
		}
	}

	// Waiting for all channel
	time.Sleep(time.Second * 1)
	assert.Equal(num, times)
}

func TestClientBroadcastAll(t *testing.T) {

	assert := assert.New(t)
	client := createClient()

	assert.NotNil(client)
	socket.Initialize()
	assert.NotNil(socket.Clients)
	assert.Equal(0, len(socket.Clients))
	socket.Clients[client.Handshake] = client

	num := 10000

	for i := 0; i < num; i++ {
		client_ := createClient()
		socket.Clients[client_.Handshake] = client_
	}

	client.Socket = socket

	assert.Equal(num + 1, len(client.Socket.Clients))

	client.Broadcast("test", Json {
		"key1" : "value1",
		"key2" : "value2",
	})

	times := 0

	for _, client_ := range socket.Clients {
		go func(client_ Client) {
			times ++
			assert.Equal(Json {
				"event" : "test",
				"data"  : Json {
					"key1" : "value1",
					"key2" : "value2",
				},
			}, <- client_.Output)
		} (client_)
	}

	// Waiting for all channel
	time.Sleep(time.Second * 1)
	assert.Equal(num + 1, times)
}