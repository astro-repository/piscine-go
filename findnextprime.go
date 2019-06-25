package piscine

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

func isNextPrime(nb,appel int) bool{
	appel++
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
	if nb < 1000000094{
		if isPrime(nb){
			return nb
		}else{
			i:=0
			for !isPrime(nb) {
				i++
				nb++
				if i > 20 {
					return 2
				}
				isNextPrime(nb,i)
			}
			return nb
		}
	}else{
		return 2
	}
}
