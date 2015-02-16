package main

func main() {

	socket := Socket {
		Port: 3000,
		Token: "acbz@3345123124567",
		Transport: Polling,
	}

	socket.Static("/", "./assert")
	socket.Handle()
}