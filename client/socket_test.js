var $ = QUnit;

$.test("get option", function(assert) {

	// Test 1
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
	this.socket = new Socket(3000);

	var option = {
		url : "/polling",
	};

	this.socket.sync(this, option, function(socket, data) {
		assert.equal(data.event, "connection");
	});
});

$.test("asynchronous request", function(assert) {
	this.socket = new Socket(3000);

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
	this.socket = new Socket(3000);

	// Handshake was created
	assert.ok(true, typeof this.socket.handshake !== 'undefined');

	// Handshake is a string
	assert.ok(true, typeof this.handshake === 'string');

	// Compare handshake length
	assert.ok(true, this.socket.handshake.length == 20);
});

$.test("process response data", function(assert) {
	this.socket = new Socket(3000);

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
	this.socket = new Socket(3000);
	var done = assert.async();

	this.socket.push(this, {
        "event" : event,
        "data"  : data
    }, function(socket, data) {
    	assert.equal(data.status, "OK");
        done();
    });
});
