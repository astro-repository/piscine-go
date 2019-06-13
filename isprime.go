package piscine

func IsPrime(nb int) bool{
	compteur := 0
	for i := 1; i <= nb; i++ {
		if nb % i == 0{
			compteur++
		}
		if compteur == 2 && i==nb {
			return true
		}
	}
	return false
}