package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args)>=3 {
		os.Args = os.Args[1:]
		a, errA := strconv.Atoi(os.Args[0])
		b, errB := strconv.Atoi(os.Args[2])
		if errA != nil || errB != nil{
			fmt.Println(0)
		}else{
			
			switch os.Args[1] {
			case "+": fmt.Println(a + b)
			case "*": fmt.Println(a * b)
			case "p": fmt.Println(a - b)
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
			}
		}
	}else{
		fmt.Println(0)
	}
}