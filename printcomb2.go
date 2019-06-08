package piscine

import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i+1; j < 99; j++ {
			if i<10 {
				if j<10 {
					fmt.Printf("0%d 0%d",i,j)
				}else{
					fmt.Printf("0%d %d",i,j)
				}
			}else{
				if j<10 {
					fmt.Printf("%d 0%d",i,j)
				}else{
					if i==98 && j==99 {
						fmt.Printf("%d %d\n",i,j)
					}
					fmt.Printf("%d %d",i,j)
				}
			}
		}
	}
}