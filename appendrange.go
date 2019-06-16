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
+	QUEST 07 : appendrange.go					+						+
+	GITHUB : github.com/cedrick777				+						+
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*/
package piscine

import (
	"fmt"
	"math"
)

func AppendRange(min, max int) []int {
	tableau := []int{}
	for i := min; i < max ; i++ {
		if len(tableau)<math.MaxInt32 {
			tableau = append(tableau,i)
		}
	}
	return tableau
}