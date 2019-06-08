package piscine

import "fmt"

func main() {
	
}

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i+1; j <= 8; j++ {
			for k := j+1; k <= 9; k++ {
				fmt.Printf("%v%v%v, ",i,j,k)
			}
		}
	}
}