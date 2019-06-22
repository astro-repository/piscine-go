/*--------------------------------------------------.
|     **      ******  ********  *******     ****    |
|  **    **  *.          **     **   **   **    **  |
|  **    **     ***.     **     ** ***    **    **  |
|  ********  *      *    **     **   **   **    **  |
|  **    **   ******     **     **   **     ****    |
|---------------------------------------.           |
|          ZONE-01 (cedrick777)         |           |
|---------------------------------------|           |
|  NOM : TOURE Ahmed Christian Cédrick  |           |
|  QUEST 09 : sortwordarr               |           |
'---------------------------------------'----------*/
package piscine

import "fmt"

func main() {
	SortWordArr([]string{"a", "A", "1", "b", "B", "2", "c", "C", "3"})
}

func SortWordArr(array []string) {
	narray := []rune{}
	sarray := []string{}

	for _, v := range array {
		a := []rune(v)
		narray = append(narray, a[0])
	}

	for i := 0; i < len(narray)-1; i++ {
		for j := 0; j < len(narray)-1; j++ {
			if narray[j] > narray[j+1]{
				intermediaire := narray[j]
				narray[j] = narray[j+1]
				narray[j+1] = intermediaire
			}
		}
	}
	
	for _, v := range narray {
		sarray = append(sarray, string(v))
	}

	fmt.Println(sarray)
}