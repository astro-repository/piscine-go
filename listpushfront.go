package main

import "fmt"

func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(link *List, data interface{}) {

	n:=&NodeL{Data:data,Next:link.Tail}

	if link.Head == nil {
		link.Head = n
		link.Tail = n
	}else{
		link.Head = n
		link.Tail = link.Head
	}

}
