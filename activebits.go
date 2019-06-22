package piscine

import "strconv"

func ActiveBits(n int) uint {

	a:= int64(n)
	nbin := strconv.FormatInt(a, 2)
	compteur := 0
	for _, v := range nbin {
		if v == '1' {
			compteur++
		}
	}
	return uint(compteur)

}
