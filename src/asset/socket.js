var Socket  = function() {

    this.sync = function(context, option, callback) {
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

    this.async = function(context, option, callback) {
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
    };

    this.process = function(data) {
        console.log(data)
    };
};

Socket.prototype  = {
    on  : function(event, fct) {
        this._events = this._events || {};
        this._events[event] = this._events[event] || [];
        this._events[event].push(fct);
    },
    remove  : function(event, fct) {
        this._events = this._events || {};
        if( event in this._events === false)  return;
        this._events[event].splice(this._events[event].indexOf(fct), 1);
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

socket = new Socket();
socket.connect();

socket.on('test2', function(data) {
    console.log(data)
})

socket.on('test', function(data) {
    socket.emit('test2', {
        data : 'Package 2 from client'
    })
});

socket.emit('init', {
    data: 'Package from client'
})
