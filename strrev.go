package piscine

import "fmt"

func StrRev(s string) string{
	secondTableau := []byte(s)
	longueur := len(s) - 1
	for i, _ := range secondTableau {
		secondTableau[i] = s[longueur]
		longueur--
	}
	return string(secondTableau)
}