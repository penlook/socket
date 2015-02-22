package main

import (
	//"container/list"
)

type Node struct {
	Id int
	Event string
	Callback func(data Json)
}

type Graph struct {
}

func (graph Graph) Children(node int) {

}

func (graph Graph) Parent(node int) {

}
