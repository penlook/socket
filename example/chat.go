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
package main

import (
	. "github.com/penlook/socket"
	"github.com/gin-gonic/gin"
	"fmt"
)

type User struct {
	Name string
}

func main() {

	socket := Socket {
		Port: 3000,
		Token: "acbz@3345123124567",
		Interval: 60,
	}

	socket.Initialize()

	// Mapping resources
	socket.Static("/static", "./")
	socket.Static("/client", "./../client")
	socket.Template("./")

	socket.Router.GET("/", func(context *gin.Context) {
		context.HTML(200, "login.html", Json {})
	})

	socket.Router.GET("/chat/", func(context *gin.Context) {
		context.Request.ParseForm()
		context.HTML(200, "chat.html", Json {
			"username" : context.Request.Form.Get("username"),
		})
	})

	list := make([] User, 0)

	socket.On("connection", func(client Client) {
		client.On("init", func(data Json) {

			_ = append(list, User {
				Name: data["username"].(string),
			})

			client.BroadcastAll("listchat", Json {
				"users" : list,
			})
		})
	})

	socket.On("disconnect", func(client Client) {
		fmt.Println("Connection is corrupt !")
	})

	socket.Listen()
}

