package piscine

func Sqrt(nb int) int {
	result := 0
	for digit := nb; digit > 0; digit-- {
		if digit*digit == nb {
			result = digit
		}
	}
	return result
}
