package piscine

import "math"

func FindNextPrime(nb int) int{
	if nb<2{
		return 2
	}else{
		for i:=2 ; i<= int( math.Sqrt(float64(nb)) ); i++{
			if nb%i==0{
				nb++
				return FindNextPrime(nb)
			}
		}
	}
	return nb
}
