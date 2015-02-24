var Socket  = function() {
    // Constructor
    function() {
        console.log('Constructor')
    }()
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