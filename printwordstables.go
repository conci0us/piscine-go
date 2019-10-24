package piscine

import "fmt"

func PrintWordsTables(table []string) {
	for str := range table {
		Printstr(str)
		z01.PrintRune('\n')
	}
}
