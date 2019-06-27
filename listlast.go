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
package main

import "fmt"

func main() {

	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")
	ListPushBack(link, "7")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))

}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(link *List, data interface{}) {

	// n:=&NodeL{Data:data,Next:link.Tail}
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


func ListLast(link *List) interface{} {

	for link.Head != nil {
		if link.Head.Next == nil {
			return link.Head.Data
		}
		link.Head = link.Head.Next
	}

	return nil
}