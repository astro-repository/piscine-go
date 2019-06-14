package piscine

func IterativePower(nb, power int) int {
	exposant := 1
	if power<0 {
		return 0
	}else{
		for i := 0; i < power; i++ {
			exposant *= nb
		}
		return exposant
	}
}