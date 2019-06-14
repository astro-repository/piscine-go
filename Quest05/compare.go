package piscine

func Compare(s, toFind string) int {
	runeTab := []rune(s)
	runeFind := []rune(toFind)
	compteur := 0
	
	if runeTab[0] != runeFind[0] {
		return -1
	}else{
		for i := 0; i < len(runeTab); i++ {

			juge := false

			if runeTab[i] == runeFind[0] {
				juge = true

				for j := 0; j < len(runeFind); j++ {
					
					if runeTab[i] == runeFind[j] {
						i++
						juge = true
					}else{
						i++
						juge = false
						break
					}
				}
			}

			if juge {
				compteur++
				i--
			}
		}
		if compteur != 0 && len(runeTab)==len(runeFind) {
			return 0
		}else if compteur != 0 {
			return 1
		}else{
			return -1
		}
	}
}