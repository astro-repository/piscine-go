package student

import "os"

func PrintProgramname() string {
	return os.Args[0]
}