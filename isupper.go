package piscine

func IsUpper(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if !(v >= 65 && v <= 90) {
			return false
		}
	}
	return true
}