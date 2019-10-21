package piscine

func reverse(s string, reversed string) string {
	if s == "" {
		return reversed
	}
	return reverse(s[1:], string(s[0])+reversed)
}

func StrRev(s string) string {
	return reverse(s, "")
}
