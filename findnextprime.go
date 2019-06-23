package piscine

func FindNextPrime(nb int) int {
	if nb < 1000000094 && 0 < nb{
		compteur := 0
		for i := 1; i <= nb; i++ {
			if nb % i == 0{
				compteur++
			}
			if compteur == 2 && i==nb {
				return nb
			}
			if compteur == 2 && i!=nb {
				break
			}
		}
		nb++
		return FindNextPrime(nb)
	}
}
