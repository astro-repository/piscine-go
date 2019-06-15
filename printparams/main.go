package main

import "os"
import "fmt"

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}