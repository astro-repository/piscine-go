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
|  QUEST 11 : listsize                  |           |
'---------------------------------------'----------*/
package piscine

func ListSize(link *List) int {
	compteur :=0

	for link.Head != nil {
		compteur++
		link.Head = link.Head.Next
	}

	return compteur
}