package main

func main()  {
	fmt.Println(ToLower("Hello!"))
	fmt.Println(ToLower("Salut!"))
	fmt.Println(ToLower("Ola!"))
}

func ToLower(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	for i, v := range sRune {
		if v >= 41 && v <= 90 {
			sRune[i] = v + distance
		}
	}
	return string(sRune)
}