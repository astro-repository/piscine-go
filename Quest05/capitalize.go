package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	k:=1
	if (sRune[0] >= 97 && sRune[0] <= 122 ) {
		sRune[0] = sRune[0] - distance
	}

	for i := 1; i < len(sRune); i++ {
		if ((sRune[i] >= 97 && sRune[i] <= 122) || (90 <= sRune[i] && sRune[i] >= 65)) && (97 <= sRune[i-1] && sRune[i-1] <= 122) && (65 <= sRune[i-1] && sRune[i-1] <= 90) {
			sRune[i] = sRune[i]
		}else if (sRune[i] >= 97 && sRune[i] <= 122) && !(97 <= sRune[i-1] && sRune[i-1] <= 122) && !(65 <= sRune[i-1] && sRune[i-1] <= 90) && !(48 <= sRune[i-1] && sRune[i-1] <= 57){
			sRune[i] = sRune[i] - distance
		}else {
			sRune[i] = sRune[i]
		}
		k++
	}
	
	return string(sRune)
}