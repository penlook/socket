package socket

type Node struct {
	Id int
	Event string
	Callback func(data Json)
}

