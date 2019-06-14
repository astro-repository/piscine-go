package piscine

func NRune(s string, n int) rune {
	runeTab := []rune(s)
	return runeTab[n-1]
}