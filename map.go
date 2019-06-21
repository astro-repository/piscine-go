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
|  QUEST 09 : Map                       |           |
'---------------------------------------'----------*/
package piscine

func Map(f func(int) bool, arr []int) []bool {
	tabRetour := []bool{}
	for _, v := range arr {
		tabRetour = append(tabRetour, f(v))
	}
	return tabRetour
}