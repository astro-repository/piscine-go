package piscine

import "fmt"

func UltimateDivMod(a, b *int) {
	*a = a/b
	*b = a%b
}