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
|  QUEST 07 : concatparams              |            |
'---------------------------------------'-----------*/
package piscine

func ConcatParams(test []string) string  {
	concatTest :=""
	for i, v:= range test {
		if i<len(test)-1 {
			concatTest += v +"\n"
		}else{
			concatTest += v
		}
	}
	return concatTest
}