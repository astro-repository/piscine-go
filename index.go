package piscine

import "math"

func Index(s, toFind string) int {

	runeTab := []rune(s)
	runeFind := []rune(toFind)
	index := -1
	
	if len(toFind) == 0{
		return 0
	}
	if len(s) <= math.MaxInt32 && len(toFind) <= len(s){

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
				index=i
				return index - len(runeFind)
			}
		}
	}
	return -1
}