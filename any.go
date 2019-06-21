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
|  QUEST 09 : Any                       |           |
'---------------------------------------'----------*/
package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, v := range arr {
		v = rune(v)
		for _, val := range v {
			if 48<=val && val <= 57 {
				return true
			}
		}
	}
	return false
}