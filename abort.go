package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}

	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1]{
				intermediaire := table[j]
				table[j] = table[j+1]
				table[j+1] = intermediaire
			}
		}
	}
	return table[2]
}