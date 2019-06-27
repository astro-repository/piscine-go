package piscine

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
