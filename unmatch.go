package piscine

func Unmatch(arr []int) int {
	arrT := Maxm(arr)

	for i := 0; i < len(arrT)-1; i++ {
		for j, v := range arrT {
			if arrT[i] == v {
				arrT[i] = 0
				arrT[j] = 0
				break
			}

			if i==len(arrT)-1 {
				return arrT[i]
			}
		}
	}

	for _, v := range arrT {
		if v != 0 {
			return v
		}
	}

	return 0
}

func Maxm(table []int) []int {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1]{
				intermediaire := table[j]
				table[j] = table[j+1]
				table[j+1] = intermediaire
			}
		}
	}
	return table
}
