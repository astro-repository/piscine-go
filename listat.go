package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func ListAt(link *NodeL, pos int) *NodeL{

	compteur :=-1

	for link.Next != nil {
		compteur++
		if compteur == pos{
			return link
		}
		link = link.Next
	}

	// if pos<0 || pos>compteur {
	// 	return nil
	// }
		Link := link{}
	return &link

}


func ListPushBack(link *List, data interface{}) {

	// n:=&NodeL{Data:data,Next:link.Tail}
	n:=&NodeL{Data:data}
	if link.Head == nil {
		link.Head = n
		// link.Tail = n
	}else{
		// link.Head = n
		// link.Tail = link.Head
		
		for i := link.Head; i != nil; i = i.Next {
			if i.Next == nil {
				i.Next = n
				break
			}
		}
	}

}