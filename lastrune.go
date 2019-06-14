package piscine


func LastRune(s string) rune {
	runeTab := []rune(s)
	return runeTab[len(runeTab)-1]
}