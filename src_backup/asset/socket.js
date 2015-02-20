var Socket  = function() {

    this.send = function(context, option, callback) {
        var request = new XMLHttpRequest();
            request.open(option.method, option.url);
            request.onreadystatechange = function () {
                if (request.readyState != 4 || request.status != 200) return;
                callback(context, JSON.parse(request.responseText));
            };
            request.send("a=1&b=2");
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
    emit : function(event /* , args... */){
        this._events = this._events || {};
        if (event in this._events === false)  return;
        for (var i = 0; i < this._events[event].length; i++) {
            this._events[event][i].apply(this, Array.prototype.slice.call(arguments, 1));
        }
    },
    connect: function() {
        var option = {
            method: "GET",
            url: "/polling"
        };

        this.send(this, option, function(socket, data) {
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
                url: "/polling?handshake=" + this.handshake,
            };

            this.send(this, option, function(socket, data) {
                socket.pull();
                socket.process(data);
            });
        }
    },
    push: function() {

        var option = {
            method: "POST",
            url: "/polling",
            data: {},
        };
    }
};

socket = new Socket();
socket.connect();

socket.on('connection', function(data) {
    socket.emit('test', {
        key: "value"
    })
});

socket.on('init', function(data) {
    console.log(data);
});
