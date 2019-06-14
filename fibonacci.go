package main

import "fmt"

func main()  {
	arg1 := 25
	fmt.Println(Fibonacci(arg1))
}

func Fibonacci(arg1 int) int{
	if arg1 == 1 || arg1 == 2 {
		return 1
	}else if arg1 < 0{
		return -1
	}else{
		return Fibonacci(arg1-2) + Fibonacci(arg1-1)
	}
}