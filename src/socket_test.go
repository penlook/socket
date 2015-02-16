package socket

import	(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSocket(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Test", "Test")

	socket := Socket {
		Port: 3000,
		Transport: Polling,
	}

	socket.Handle()
	socket.Emit("message", Json {
		"data" : "abc",
		"test" : "123",
	})
}


