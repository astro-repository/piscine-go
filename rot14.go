package main

import "fmt"

func main() {
	
	fmt.Println(Rot14("V 8G8opSMhg41"))
	
}

func Rot14(str string) string {
	Maj := alphabetMaj()
	Min := alphabetMin()
	strrune := []rune(str)
	for _, v := range strrune {
		if 'A' <= v && v <= 'Z' {
			index := returnIndex(v) + 1
			sliceMaj := Maj[index:]
			sliceMajRestant := Maj[:index]
			hachage := append(sliceMaj,sliceMajRestant...)
			fmt.Print(string(hachage[13]))
		}
		
		if 'a' <= v && v <= 'z' {
			index := returnIndex(v) + 1
			sliceMin := Min[index:]
			sliceMinRestant := Min[:index]
			hach := append(sliceMin,sliceMinRestant...)
			fmt.Print(string(hach[13]))
		}

		if !('A' <= v && v <= 'Z') && !('a' <= v && v <= 'z') {
			fmt.Printf(string(v))
		}

	}
	return ""
}

func alphabetMaj() []rune  {
	Maj := []rune{}
	for i := 'A'; i <= 'Z'; i++ {
		Maj = append(Maj, i)
	}
	return Maj
}

func alphabetMin() []rune  {
	Min := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		Min = append(Min, i)
	}
	return Min
}

func returnIndex(car rune) int {
	for i, v := range alphabetMaj() {
		if car == v {
			return i
		}
	}

	for i, v := range alphabetMin() {
		if car == v {
			return i
		}
	}
	return -1
}