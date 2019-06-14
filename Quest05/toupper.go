package main

func main()  {
	fmt.Println(ToUpper("Hello!"))
	fmt.Println(ToUpper("Salut!"))
	fmt.Println(ToUpper("Ola!"))
}

func ToUpper(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	for i, v := range sRune {
		if v >= 97 && v <= 122 {
			sRune[i] = v - distance
		}
	}
	return string(sRune)
}