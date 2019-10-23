package piscine

func TrimAtoi(s string) int {
	result := 0
	minus := false
	for _, char := range s {
		if char == '-' && result == 0 {
			minus = true
		}
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-48)
		}
	}
	if minus {
		return -result
	} else {
		return result
	}
}
