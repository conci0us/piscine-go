package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	if str != "" {
		z01.PrintRune(rune(str[0]))
		PrintStr(str[1:])
	}
}
