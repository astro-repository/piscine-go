package piscine

import "fmt"

func PrintNbrBase(nbr int, base string) {
	if nbr<0 {
		fmt.Printf("-")
		nbr = -nbr
	}
	
	a := pVerification{
		nbr : nbr,
		base : []rune(base),
	}
	tab:=[]int{}

	if verification(&a) {
		tab = append(tab, a.nbr % len(a.base))
		for int(a.nbr / len(a.base))> 0 {
			a.nbr = int(a.nbr / len(a.base))
			tab = append(tab, a.nbr % len(a.base))
		}

		for i := len(tab)-1; 0 <= i; i-- {
			fmt.Printf(string(base[tab[i]]))
		}
	}else{
		fmt.Printf("NV")
	}
}

type pVerification struct {
	nbr int
	base []rune
}

func verification(a *pVerification) bool {
	if len(a.base) >= 2 {
		for i, V := range a.base {
			if V == '+' || V == '-' {
				return false
			}
			for j := i+1; j < len(a.base); j++ {
				if V == a.base[j] {
					return false
				}
			}
		}
	}
	return true
}
