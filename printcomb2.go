package piscine

import "fmt"

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for s := 0; s <= 8; s++ {
			for t := a; t <= 9; t++ {
				for r := 0; r <= 9; r++ {
					if a==9 && s==8 && t==9 && r==9 {
						fmt.Printf("%d%d %d%d\n",a,s,t,r)
					}else{
						fmt.Printf("%d%d %d%d, ",a,s,t,r)
					}
				}
			}
		}
	}
}