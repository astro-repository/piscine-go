package main

import "github.com/01-edu/z01"
import "os"

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr%2) == 1 {
		return true
	} else {
		return false
	}
}

func even(nbr int) int {
	if even(nbr % 2) == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lengthOfArg := len(os.Args[1:len(os.Args)])

	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}