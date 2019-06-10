package piscine

import "fmt"

func main() {
	var index rune
	index = 97
	for index <= 122 {
		fmt.Printf("%c", index)
		index++
	}
	print("\n")
}