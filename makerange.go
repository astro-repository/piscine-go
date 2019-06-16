/*++++++++++++++++++++++++++++++++++++++++++++++++++++++
+                                                      +
+     **      ******   ********  *********     ****    +
+  **    **  **     *  ** ** **  **     **   **    **  +
+  **    **  **           **     **     **   **    **  +
+  ********     ** **     **     ** ****     **    **  +
+  **    **  *      *     **     **     **   **    **  +
+  **    **   ******      **     **     **     ****    +
+                                                      +
++++++++++++++++++++++++++++++++++++++++++++++++++++++++
+              ZONE-01               +                 +
++++++++++++++++++++++++++++++++++++++                 +
+  NOM : TOURE                       +                 +
+  PRENOMS : Ahmed Christian Cédrick +                 +
+  QUEST 07 : makerange              +                 +
+  GITHUB : github.com/cedrick777    +                 +
++++++++++++++++++++++++++++++++++++++++++++++++++++++*/
package piscine

import (
	"math"
)

func MakeRange(min, max int) []int {
	if max - min <= 0 {
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