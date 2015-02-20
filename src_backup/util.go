package main

import(
	"crypto/rand"
)

func random() string {
    dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, 20)
    rand.Read(bytes)
    for k, v := range bytes {
        bytes[k] = dictionary[v%byte(len(dictionary))]
    }
    return string(bytes)
}