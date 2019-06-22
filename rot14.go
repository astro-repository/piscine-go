package piscine

import "fmt"

func Rot14(str string) string {
	Maj := alphabetMaj()
	Min := alphabetMin()
	strrune := []rune(str)
	chaineF := []rune{}

	for _, v := range strrune {
		if 'A' <= v && v <= 'Z' {
			index := returnIndex(v) + 1
			sliceMaj := Maj[index:]
			sliceMajRestant := Maj[:index]
			hachage := append(sliceMaj,sliceMajRestant...)
			chaineF = append(chaineF, hachage[13])
		}
		
		if 'a' <= v && v <= 'z' {
			index := returnIndex(v) + 1
			sliceMin := Min[index:]
			sliceMinRestant := Min[:index]
			hach := append(sliceMin,sliceMinRestant...)
			chaineF = append(chaineF, hach[13])
		}

		if !('A' <= v && v <= 'Z') && !('a' <= v && v <= 'z') {
			chaineF = append(chaineF, v)
		}

	}
	return string(chaineF)
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