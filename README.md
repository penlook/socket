# socket
Real-time library for Go based on Long-Polling

# Documentation
Server
```go
socket := Socket {
	Port: 3000,
	Token: "acbz@3345123124567",
	Transport: LongPolling,
	Template: "asset/*",
}

socket.Initialize()
socket.Static("/static", "./asset")

socket.On("connection", func(client Client) {
	client.On("init", func(data Json) {
	})
})
```

Client
```javascript
socket = new Socket();
socket.connect();

socket.on('test', function(data) {
	socket.emit('test2', {
    	data : 'Package 2 from client'
	})
});

socket.emit('init', {
	data: 'Package from client'
})
```
