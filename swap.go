package piscine

func Swap(a *int, b *int) {
	copyA := *a
	*a = *b
	*b = copyA
}
