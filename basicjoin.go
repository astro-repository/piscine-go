package piscine

func BasicJoin(s []string) string {
	concatenation := ""
	for _, v := range s {
		concatenation += v
	}
	return concatenation
}