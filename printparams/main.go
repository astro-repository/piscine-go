package main

import "os"
import "fmt"

func main() {
	for _, v := range os.Args {
		fmt.Println(v)
	}
}