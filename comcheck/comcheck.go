package main

import "fmt"
import "os"

func main() {
	if len(os.Args)>1 {
		os.Args = os.Args[1:]
		juge := false
		for _, v := range os.Args {
			if v == "01" || v == "galaxy" || v == "galaxy 01" {
				juge = true
			}
		}
		if juge {
			fmt.Println("Alert!!!")
		}
	}
}

