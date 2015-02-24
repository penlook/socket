package socket

import(
	"crypto/rand"
)

// JSON data
type Json map[string] interface {}

// Event structure
type Event struct {
	Id int
	Name string
	Callback func(data Json)
}

func random() string {
    dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, 20)
    rand.Read(bytes)
    for k, v := range bytes {
        bytes[k] = dictionary[v%byte(len(dictionary))]
    }
    return string(bytes)
}