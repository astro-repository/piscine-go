package piscine

import (
	"fmt"
	student "./student"
)

func Raid1a(x,y int){
	for i := 1; i <= y; i++ {
		if x==1 && y!=1 {
		for j := 1; j <= y; j++ {

			if j==1 || j==y {
				print("o\n")
			}else if j!=1 && j!=y {
				print("|\n")
			}
		}
		break
		}else{
			for j := 1; j <= x; j++ {
				if (i==1 && j==1) || (i==1&&j==x) || (i==y && j==1) || (i==y && j==x)  {
					print("o")
				}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
					print("-")
				}else if (i!=1 && j==1) || (i!=y && j==x){
					print("|")
				}else{
					print(" ")
				}
			}
		}
		print("\n")
	}
}