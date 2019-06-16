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
|  QUEST 07 : SplitWhiteSpaces          |            |
'---------------------------------------'-----------*/
package piscine

func SplitWhiteSpaces(texte string) []string {
	TexteToString := ""
	tableau := []string{}
	for i, v := range texte {
		if i == len(texte)-1 && string(v) != " " && string(v) != "	" && string(v) != "\n" {
			TexteToString += string(v)
			tableau = append(tableau, TexteToString)
		}else if string(v) != " " && string(v) != "	" && string(v) != "\n" {
			TexteToString += string(v)
		}else{
			if len(TexteToString)>=1{
				tableau = append(tableau, TexteToString)
			}
			TexteToString = ""
		}
	}
	return tableau
}