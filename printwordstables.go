package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, str := range table {
		Printstr(str)
		z01.PrintRune('\n')
	}
}
