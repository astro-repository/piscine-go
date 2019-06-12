package piscine

import "fmt"
import "math"

func BasicAtoi2(s string) int{
	NombreIdentique := []byte(s)
	for i, _ := range NombreIdentique {
		switch NombreIdentique[i] {
			case 48: NombreIdentique[i] = byte(0)
			case 49: NombreIdentique[i] = byte(1)
			case 50: NombreIdentique[i] = byte(2)
			case 51: NombreIdentique[i] = byte(3)
			case 52: NombreIdentique[i] = byte(4)
			case 53: NombreIdentique[i] = byte(5)
			case 54: NombreIdentique[i] = byte(6)
			case 55: NombreIdentique[i] = byte(7)
			case 56: NombreIdentique[i] = byte(8)
			case 57: NombreIdentique[i] = byte(9)
			default :
			if NombreIdentique[i]!=57{
				return int(0)
			}
		}
	
	}
	entier := 0
	longueur := len(NombreIdentique) - 1
	for j := 0; j < len(NombreIdentique); j++ {
		entier += int(NombreIdentique[j]) * int(math.Pow10(longueur))
		longueur--
	}
	return entier
}