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
	//"github.com/gin-gonic/gin"
)

func TestSocket(t *testing.T) {

	assert := assert.New(t)
	assert.Equal("Test", "Test")

	// Implemting .. Mockup HTTP Request
	
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


