package piscine

func FindNextPrime(nb int) int{
	compteur := 0
	for i := 1; i <= nb; i++ {
		if nb % i == 0{
			compteur++
		}
		if compteur == 2 && i==nb {
			return nb
		}
		if compteur == 2 && i!=nb{
			return FindNextPrime(nb+1)
		}
	}
	return 0
}