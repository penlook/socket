package main

type Node struct {
	Id int
	Event string
	Callback func(client Client)
}

type Vertex struct {
	X int16
	Y int16
}

type Graph struct {
	Vertex [] Vertex
}

func (graph Graph) Children(node int) {

}

func (graph Graph) Parent(node int) {

}
