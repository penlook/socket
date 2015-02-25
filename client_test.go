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
	//"fmt"
)

func createClient() (Client, string) {

	socket := Socket {
		Port : 3000,
		Interval: 60,
	}

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

	return client, handshake
}

func TestClientOn(t *testing.T) {

	assert := assert.New(t)
	client, _ := createClient()

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
	client, _ := createClient()

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
	client, _ := createClient()

	assert.NotNil(client)
	client.Socket.Initialize()
	assert.NotNil(client.Socket.Clients)
	assert.Equal(0, len(client.Socket.Clients))

	for i := 0; i < 10000; i++ {
		client_, handshake := createClient()
		client.Socket.Clients[handshake] = client_
	}

	assert.Equal(10000, len(client.Socket.Clients))

	client.Broadcast("test", Json {
		"key1" : "value1",
		"key2" : "value2",
	})

	for _, client_ := range client.Socket.Clients {
		assert.Equal(Json {
			"event" : "test",
			"data"  : Json {
				"key1" : "value1",
				"key2" : "value2",
			},
		}, <- client_.Output)
	}
}
