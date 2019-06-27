package piscine

func ListAt(link *NodeL, pos int) *NodeL{
	compteur :=-1

	for link != nil {
		compteur++
		if compteur == pos{
			return link
		}
		link = link.Next
	}
	return link
}
