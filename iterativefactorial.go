package piscine

import "math"

func IterativeFactorial(nb int) int {
		factoriel := 1
		for i := 1; i <= nb; i++ {
			if factoriel > math.MaxInt64 || nb < 0 || factoriel < 0{
				return 0
			}else if nb == 0 {
				return factoriel
			}else{
				factoriel *= i
			}
		}
		if nb<0{
			return 0
		}
		return factoriel
}