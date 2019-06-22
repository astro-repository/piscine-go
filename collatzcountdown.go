package piscine

func CollatzCountdown(start int) int {
	compteur := 1
	return Compteur(compteur, start)
}

func Compteur(compteur int, start int) int {

	if start == 1{
		return compteur
	}
	
	if start % 2 == 0 {
		start = start / 2
		compteur++
		return Compteur(compteur, start)
	}else{
		start = (start * 3)+1
		compteur++
		return Compteur(compteur, start)
	}

	return compteur
}