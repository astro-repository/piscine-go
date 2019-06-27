package piscine

func ListClear(link *List) {
	for link.Head != nil {
		link.Head = nil
	}
}
