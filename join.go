package piscine

func Join(s []string, separateur string) string {
	concatenation := ""
	for i, v := range s {
		if i != len(s)-1 {
			concatenation += v + separateur
		}else{
			concatenation += v
		}
	}
	return concatenation
}