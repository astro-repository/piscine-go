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
|  QUEST 11 : listlast                  |           |
'---------------------------------------'----------*/
package piscine

func ListLast(link *List) interface{} {

	for link.Head != nil {
		if link.Head.Next == nil {
			return link.Head.Data
		}
		link.Head = link.Head.Next
	}

	return nil
}
