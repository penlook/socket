var socket = new Socket(3000);

socket.on('test_on', function(data) {
	console.log(data)
})

socket.on('join', function(data) {
	console.log(data)
})

socket.on('listchat', function(data) {
	console.log(data)
})

socket.emit('init', {
    username: getParamByName('username'),
});