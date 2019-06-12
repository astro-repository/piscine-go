package piscine

func UltimateDivMod(a, b *int) {
	A = *a
	*a = *a / *b
	*b = A % *b
}