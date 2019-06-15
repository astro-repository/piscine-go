package piscine

func IsPrintable(s string) bool {
	sRune := []rune(s)
	for _, v := range sRune {
		if v==92 && (v >= 97 && v <= 122) {
			return false
		}
	}
	return true
}