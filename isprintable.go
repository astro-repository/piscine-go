package piscine

func IsPrintable(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if !(v >= 90 && v <= 65) {
			return true
		}
	}
	return true
}