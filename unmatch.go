package piscine

func Unmatch(arr []int) int {
	table := make([]int, len(arr))
	
	for i := 0; i < len(arr) - 1; i++ {
		compteur := 0
		for j := 1; j < len(arr) - 1; j++ {
			if arr[i] == arr[j] {
				compteur++
			}
		}
		table[i] = compteur
	}

	for i, v := range table {
		if v == 0 {
			return arr[i]
		}
	}
	return 0
}
