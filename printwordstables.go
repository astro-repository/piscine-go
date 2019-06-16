/*---------------------------------------------------.
|     **      ******   ********  *******     ****    |
|  **    **  *.           **     **   **   **    **  |
|  **    **     ***.      **     ** ***    **    **  |
|  ********  *      *     **     **   **   **    **  |
|  **    **   ******      **     **   **     ****    |
|---------------------------------------.            |
|          ZONE-01 (cedrick777)         |            |
|---------------------------------------|            |
|  NOM : TOURE Ahmed Christian Cédrick  |            |
|  QUEST 07 : printwordstables          |            |
'---------------------------------------'-----------*/
package piscine

func main()  {
	PrintWordsTables([]string{"Hello", "how", "are", "you?"})
}
func PrintWordsTables(table []string) {
	for _, v := range table {
		fmt.Println(v)
	}
}