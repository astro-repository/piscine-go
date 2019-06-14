package piscine

func ToUpper(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	for i, v := range sRune {
		if v >= 97 && v <= 122 {
			sRune[i] = v - distance
		}
	}
	return string(sRune)
}