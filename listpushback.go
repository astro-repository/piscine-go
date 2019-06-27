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
|  QUEST 11 : listpushback              |           |
'---------------------------------------'----------*/
package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(link *List, data interface{}) {

	n:=&NodeL{Data:data}
	if link.Head == nil {
		link.Head = n
	}else{
		for i := link.Head; i != nil; i = i.Next {
			if i.Next == nil {
				i.Next = n
				break
			}
		}
	}

}
