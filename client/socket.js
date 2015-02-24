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

var Socket  = function() {
    // Constructor
    console.log('Constructor')
};

Socket.prototype  = {

    // Synchronous request
    sync : function(context, option, callback) {
        var request = new XMLHttpRequest();
        request.open(option.method, option.url, false);
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.send(option.data);

        try {
            data = JSON.parse(request.responseText);
        } catch (e) {
            console.log(e);
            return false;
        }

        console.log(data);

        callback(context, data);
    },

    // Asynchronous request
    async : function(context, option, callback) {
        var request = new XMLHttpRequest();
        request.open(option.method, option.url);
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.onreadystatechange = function () {
            if (request.readyState != 4 || request.status != 200) return;

            var data = {}
            try {
                data = JSON.parse(request.responseText);
            } catch (e) {
                console.log(e);
                return false;
            }

            option.async = false;
            callback(context, data);
        };
        request.send(JSON.stringify(option.data));
    },

    // Processor
    process : function(data) {
        console.log(data)
    },

    // Register event
    on : function(event, callback) {

    },

    // Remove event
    remove : function(event) {

    },

    emit : function(event, data) {
        console.log("Emit")
        this.push(this, {
            "event" : event,
            "data"  : data
        }, function(socket, data) {
            console.log("Emitted !");
            console.log(data);
        });
    },

    // Establish new connection
    connect: function() {
        var option = {
            method: "GET",
            url: "/polling",
            data: {},
        };
        console.log("send connection")
        this.sync(this, option, function(socket, data) {
            if (data.event == "connection") {
                socket.handshake = data.data.handshake;
                socket.pull();
            }
        });
    },

    // Pull data by using polling request
    pull: function() {
        if (typeof this.handshake === 'string') {
            var option = {
                method: "GET",
                url: "/polling/" + this.handshake,
                data: {}
            };

            this.async(this, option, function(socket, data) {
                socket.pull();
                socket.process(data);
            });
        }
    },

    // Send data to server
    push: function(context, data, callback) {
        console.log("Push to server")
        if (typeof this.handshake === 'string') {

            var option = {
                method: "POST",
                url: "/polling/" + this.handshake,
                data: data
            };

            console.log("send connection")
            this.async(this, option, function(socket, data) {
                console.log("Push done")
                //callback(socket, data)
            });
        }
    }
};