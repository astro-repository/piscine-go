package piscine

func ListReverse(link *List) {

	ElmtCourant := link.Head
	var ElmtPrecedent *NodeL
	link.Tail = link.Head

	for ElmtCourant != nil {
		next := ElmtCourant.Next
		ElmtCourant.Next = ElmtPrecedent
		ElmtPrecedent = ElmtCourant
		ElmtCourant = next
	}

	link.Head = ElmtPrecedent
}