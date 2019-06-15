package main

import "os"
import "fmt"

func main() {
	for i := len(os.Args)-1; 0 < i; i-- {
		fmt.Println(os.Args[i])
	}
}