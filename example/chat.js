var socket = new Socket(3000);

socket.on('test_on', function(data) {
	console.log(data)
})

var arr = document.URL.match(/username=([A-Za-z0-9]+)/)
var username = arr[1];
console.log(username);

socket.on('join', function(data) {
	console.log(data)
})

socket.on('listchat', function(data) {
	console.log(data)
})

socket.emit('init', {
    username: username
});