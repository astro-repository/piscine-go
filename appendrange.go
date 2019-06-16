/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+																		+
+		**		**********	* ****** *	**********	   *****			+
+	**		**	**		**	*** ** ***	**	**	**	***		***			+
+	**		**	***				**		**		**	***		***			+
+	** ****	**		** ***		**		**	**		***		***			+
+	**		**			**		**		**		**	***		***			+
+	**		**	**********		**		**		**	   *****			+
+																		+
+ +++++++++++++++++++++++++++++++++++++++++++++	+						+
+					ZONE-01						+						+
+ +++++++++++++++++++++++++++++++++++++++++++++	+						+
+	NOM : TOURE 								+						+
+	PRENOMS : Ahmed Christian Cédrick			+						+
+	QUEST 06 : sortparams						+						+
+	GITHUB : github.com/cedrick777				+						+
+ +++++++++++++++++++++++++++++++++++++++++++++ +++++++++++++++++++++++*/

package main

import "fmt"

func main() {
	fmt.Println(AppendRange(4900,5000))
	fmt.Println(AppendRange(10, 5))
	fmt.Println(AppendRange(0, 0))
}

func AppendRange(min, max int) []int {
	if min>max {
		return []int{}
	}
	tableau := make([]int, max-min)
	for i := min; i < max ; i++ {
		tableau = append(tableau,i,)
	}
	return tableau
}