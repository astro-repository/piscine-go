package piscine

import "fmt"

func PrintStr(a string) {
	// a := []string
	for _, v := range a {
		fmt.Print(string(v))
	}
	fmt.Print("\n")
}