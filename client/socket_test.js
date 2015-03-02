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

var $ = QUnit;
var socket = function() {
	return new Socket("52.10.223.209", 3000);
};

$.test("get option", function(assert) {

	// Test 1
	var custom_option = {
		key2 : "value3",
		key3 : "value1",
	};

	var default_option = {
		key1 : "value1",
		key2 : "value2",
		key3 : "value3",
		key4 : "value4",
	};

	var output = getOption(custom_option, default_option);

	assert.deepEqual(output, {
		key1 : "value1",
		key2 : "value3",
		key3 : "value1",
		key4 : "value4",
	});

	// Test 2
	custom_option = {
	};

	default_option = {
		key1 : "value1",
		key2 : "value2",
		key3 : "value3",
		key4 : "value4",
	};

	output = getOption(custom_option, default_option);

	assert.deepEqual(output, {
		key1 : "value1",
		key2 : "value2",
		key3 : "value3",
		key4 : "value4",
	});

	// Test 3
	custom_option = {
		key1 : "value1",
		key2 : "value2",
		key3 : "value3",
		key4 : "value4",
	};

	default_option = {
	};

	output = getOption(custom_option, default_option);

	assert.deepEqual(output, {
		key1 : "value1",
		key2 : "value2",
		key3 : "value3",
		key4 : "value4",
	});
});

$.test("synchronous request", function(assert) {
	this.socket = socket();

	var option = {
		url : "/polling",
	};

	this.socket.sync(this, option, function(socket, data) {
		assert.equal(data.event, "connection");
	});

});

$.test("asynchronous request", function(assert) {
	this.socket = socket();

	// Asynchronous wating ...
	var done = assert.async();

	var option = {
		url : "/polling",
	};

	this.socket.async(this, option, function(socket, data) {
		assert.equal(data.event, "connection");
		done();
	});
});

$.test("create handshake", function(assert) {
	this.socket = socket();

	// Handshake was created
	assert.ok(true, typeof this.socket.handshake !== 'undefined');

	// Handshake is a string
	assert.ok(true, typeof this.handshake === 'string');

	// Compare handshake length
	assert.ok(true, this.socket.handshake.length == 20);
});

$.test("process response data", function(assert) {
	this.socket = socket();

	var data = {
		key1: "value1",
		key2: "value2",
		key3: "value3"
	}

	// Register test event
	this.socket.on("test", function(result) {
		assert.deepEqual(result, data);
	})

	var raw_data = {
		"event" : "test",
		"data"  : data
 	}

 	// Process mock data
	this.socket.process(raw_data);
});

$.test("push data to server", function(assert) {
	this.socket = socket();
	var done = assert.async();

	this.socket.push(this, {
        "event" : event,
        "data"  : data
    }, function(socket, data) {
    	assert.equal(data.status, "OK");
        done();
    });
});