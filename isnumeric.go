package piscine

func IsAlpha(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if !(v >= 48 && v <= 57) {
			return false
		}
	}
	return true
}