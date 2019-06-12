package piscine

func Swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}