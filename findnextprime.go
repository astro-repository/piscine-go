package piscine

import "math"

func isPrime(nb int) bool{
	compteur := 0
		for i := 1; i <= nb; i++ {
			if nb % i == 0{
				compteur++
			}
			if compteur == 2 && i==nb {
				return true
			}
			if compteur == 2 && i!=nb {
				return false
			}
	}
	return false
}

func FindNextPrime(nb int) int {
	if nb < math.MaxInt32{
		if isPrime(nb){
			return nb
		}else{
			for !isPrime(nb) && (nb <= math.MaxInt32) {
				nb++
				FindNextPrime(nb)
			}
			return nb
		}
	}else{
		return 0
	}
}
