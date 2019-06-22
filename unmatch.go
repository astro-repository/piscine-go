package piscine

func Unmatch(arr []int) int {
	table := make([]int, len(arr))
	
	for i := 0; i < len(arr); i++ {
		compteur := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				compteur++
			}
		}
		table[i] = compteur
	}
	for i, v := range table {
		if v == 1 {
			return arr[i]
		}
	}
	return 0
}
