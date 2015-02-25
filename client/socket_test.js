var $ = QUnit;

$.module("test socket", {

	// Setup
	beforeEach: function() {
    	this.socket = new Socket(3000);
  	},

  	// Teardown
  	afterEach: function() {
  		this.socket.removeAllEvent();
  	}
});

$.test("get options", function(assert) {

	custom_option = {
		key2 : "value3",
		key3 : "value1",
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
		key2 : "value3",
		key3 : "value1",
		key4 : "value4",
	});

});

$.test("synchronous request", function(assert) {
	assert.equal("test", "test");
});

$.test("asynchronous request", function(assert) {
	assert.equal("test", "test");
});

$.test("create handshake", function(assert) {

	// Handshake was created
	assert.ok(true, typeof this.socket.handshake !== 'undefined');

	// Handshake is a string
	assert.ok(true, typeof this.handshake === 'string');

	// Compare handshake length
	assert.ok(true, this.socket.handshake.length == 20);
});

$.test("process response data", function(assert) {

	data = {
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

	this.socket.push(this, {
        "event" : event,
        "data"  : data
    }, function(socket, data) {
        console.log(data);
    });

    assert.equal("OK", "OK");
});
