package main

import "fmt"

func main()  {
	fmt.Println(ActiveBits(-5))
}

func ActiveBits(n int) uint {
	if n<0 {
		return uint(-n)
	}
	return uint(n)
}
