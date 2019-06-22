package piscine

func Compact(ptr *[]string, length int) int {
	compteur := 0
	for _, v := range *ptr {
		if v != " " {
			compteur++
		}
	}
	return compteur
}