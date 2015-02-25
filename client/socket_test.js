var $ = QUnit;

// Setup
$.module("test socket", {
	beforeEach: function() {
    	this.socket = Socket(3000);
  	},
});

$.test("socket ready", function() {

});

$.test("socket initialize", function(assert) {
	var socket = new Socket(300)
});

$.test("process raw data", function(assert) {
	var socket = new Socket(3000);

	data = {
		key1: "value1",
		key2: "value2",
		key3: "value3"
	}

	socket.on("test", function(result) {
		assert.deepEqual(result, data);
	})

	var raw_data = {
		"event" : "test",
		"data"  : data
 	}

	socket.process(raw_data);
});