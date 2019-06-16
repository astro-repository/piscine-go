func AppendRange(min, max int) []int {
	tableau := []int{};
	for i := min; i < max ; i++ {
		tableau = append(tableau,i);
	}
	if min>=max {
		return nil
	}
	return tableau;
}