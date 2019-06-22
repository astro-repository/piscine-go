package main

import "fmt"

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		fmt.Printf(string(s))
	}
	fmt.Println()
}



func CloseDoor(ptrDoor *Door) {
	
	PrintStr("Door Closing...")
	ptrDoor.state = false
	// state = CLOSE
	// return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("Door is open ?")
	if door.state {
		return true
	}else{
		return false
	}
	// const OPEN := "open" 
	// PrintStr("is the Door opened ?")
	// return Door.state = OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Door is closed ?")
	if ptrDoor.state {
		return false
	}else{
		return true
	}
}

type Door struct {
	state bool
}

func OpenDoor(door *Door){
	PrintStr("Door Opening...")
	door.state = true
}

// func CloseDoor(door *Door) {
// 	door.state = false
// }

func main() {
	const OPEN = true

	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
