/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+																		+
+		**		**********	* ****** *	**********	   *****			+
+	**		**	**		**	*** ** ***	**	**	**	***		***			+
+	**		**	***				**		**		**	***		***			+
+	** ****	**		** ***		**		** ****		***		***			+
+	**		**			**		**		**		**	***		***			+
+	**		**	**********		**		**		**	   *****			+
+																		+
+++++++++++++++++++++++++++++++++++++++++++++++++						+
+					ZONE-01						+						+
+++++++++++++++++++++++++++++++++++++++++++++++++						+
+	NOM : TOURE 								+						+
+	PRENOMS : Ahmed Christian Cédrick			+						+
+	QUEST 06 : sortparams						+						+
+	GITHUB : github.com/cedrick777				+						+
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*/
package piscine

import (
	"fmt"
	"math"
)

func AppendRange(min, max int) []int {
	if max - min < 0 {
		return nil
	}
	tableau := make([]int, max-min)
	for i, j := min, 0; j < len(tableau) ; i, j = i+1, j+1 {
		if len(tableau)<math.MaxInt32 {
			tableau[j] = i
		}
	}
	return tableau
}