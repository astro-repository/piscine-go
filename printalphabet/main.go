package piscine

import "github.com/01-edu/z01"

func main() {
	var index rune
	index = 97
	for index <= 122 {
		z01.PrintRune(index)
		index++
	}
	print("\n")
}