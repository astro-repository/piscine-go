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
|  QUEST 11 : listpushback              |           |
'---------------------------------------'----------*/
package main

func CountIf(f func(string) bool, tab []string) int {
	compteur := 0
	for _, v := range tab {
		if f(v) {
			compteur++
		}
	}
	return compteur
}