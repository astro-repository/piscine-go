package piscine

import "fmt"

func main()  {
	s:= []int{5,4,3,9,10,2,1,-1,0}
	SortIntegerTable(s)
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1]{
				intermediaire := table[j]
				table[j] = table[j+1]
				table[j+1] = intermediaire
			}
		}
	}
	fmt.Println(table)
}