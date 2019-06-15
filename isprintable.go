package piscine

func IsPrintable(s string) bool {
	sRune := []rune(s)

	for _, v := range sRune {
		if v >= 0 && v <= 31 {
			return false
		}
	}
	
	return true
}