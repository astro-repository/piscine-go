package piscine

func Concat(s []string) string {
	concatenation := ""
	for _, v := range s {
		concatenation += v
	}
	return concatenation
}