package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args)==4 {
		os.Args = os.Args[1:]
		a, errA := strconv.Atoi(os.Args[0])
		b, errB := strconv.Atoi(os.Args[2])
		if errA != nil {
			a := 0
		}
		if errB != nil {
			b := 0
		}

		switch os.Args[1] {
		case "+": fmt.Println(a + b)
		case "*": fmt.Println(a * b)
		case "-": fmt.Println(a - b)
		case "/": 
			if b != 0 {
				fmt.Println(a/b)
			}else{
				fmt.Println("No division by 0")
			}
		case "%": 
			if b != 0 {
				fmt.Println(a%b)
			}else{
				fmt.Println("No modulo by 0")
			}
		default :
		fmt.Println(0)
		}
	}
}
