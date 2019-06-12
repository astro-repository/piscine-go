package piscine

import "fmt"

func Swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}