package piscine

func AppendRange(min, max int) []int {
	tableau := []int{}
	for i := min; i < max ; i++ {
		tableau = append(tableau,i)
	}
	return tableau
}