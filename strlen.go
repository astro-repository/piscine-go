package piscine

func StrLen(str string) int{
	nombre := []byte(str)
	compteur := 0
	double := 1
	for _, v := range nombre {
		if v <= 127{
			compteur++
		}else{
			if double<2 {
				compteur++
				double++
			}else{
				double=1
			}
		}
	}
	return compteur
}