package main

import "fmt"

func main() {
	var index rune
	index = 122
	for 97 <= index {
		fmt.Printf("%c", index)
		index--
	}
	print("\n")
}