package piscine

func Max(table []int) int {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1]{
				intermediaire := table[j]
				table[j] = table[j+1]
				table[j+1] = intermediaire
			}
		}
	}
	return table[len(table)-1]
}