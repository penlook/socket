var socket = new Socket(3000);

socket.on('test_on', function(data) {
	console.log(data)
})

socket.on('test2', function(data) {
    console.log(data)
});

socket.on('test', function(data) {
    socket.emit('test2', {
        data : 'Package 2 from client'
    });
});

socket.emit('init', {
    data: 'Package from client'
});

console.log(socket.events)
