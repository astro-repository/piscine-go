package piscine

func ToLower(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	for i, v := range sRune {
		if v >= 65 && v <= 90 {
			sRune[i] = v + distance
		}
	}
	return string(sRune)
}