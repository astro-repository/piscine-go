package piscine

func IsAlpha(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if !(v >= 97 && v <= 122) && !(90 >= v && v >= 65) {
			return false
		}
	}
	return true
}