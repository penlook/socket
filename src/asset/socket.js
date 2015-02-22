var Socket  = function() {

    this.send = function(context, option, callback) {
        console.log(option)
        var data = option.data || {}
        var async = option.async || true

        var request = new XMLHttpRequest();
            request.open(option.method, option.url, false);
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

                callback(context, data);
            };
            console.log("Send " + JSON.stringify(option.data));
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
            async: false
        };
        console.log("send connection")
        this.send(this, option, function(socket, data) {
            if (data.event == "connection") {
                socket.handshake = data.data.handshake;
                //socket.pull();
            }
        });
    },
    pull: function() {
        if (typeof this.handshake === 'string') {
            var option = {
                method: "GET",
                url: "/polling/" + this.handshake,
                async: true
            };

            this.send(this, option, function(socket, data) {
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
                data: data,
                async: true
            };

            console.log("send connection")
            this.send(this, option, function(socket, data) {
                console.log("Push done")
                //callback(socket, data)
            });
        }
    }
};

socket = new Socket();
socket.connect();

socket.on('test', function(data) {
    console.log(data);
});

socket.emit('test', {
    key: 'value'
})