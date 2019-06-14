package piscine

func IsAlpha(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if !(v >= 97 && v <= 122) {
			return false
		}
	}
	return true
}