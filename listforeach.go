package piscine

func ListForEach(l *List, f func(*NodeL)) {

	l.Tail = l.Head
	for l.Head != nil {
		next := l.Head.Next
		f(l.Head)
		l.Head = next
	}
	l.Head = l.Tail

}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
