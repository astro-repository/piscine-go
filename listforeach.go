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