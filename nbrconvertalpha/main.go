package main

import (
	"github.com/01-edu/z01"
	"os"
)

func customStrLen(arr []string) int {
	length := 0
	for i := range arr {
		length++
		i++
	}
	return length
}

func recursiveAtoi(s string, number int) int {
	if s == "" {
		return number
	}
	return recursiveAtoi(s[1:], number*10+int(rune(s[0])-'0'))
}

func customAtoi(s string) int {
	return recursiveAtoi(s, 0)
}

func main() {
	args := os.Args
	if customStrLen(args) > 1 {
		upper := args[1]
		var numbers []string = args[1:]
		for i := 0; i < customStrLen(numbers); i++ {
			if customAtoi(numbers[i]) >= 1 && customAtoi(numbers[i]) <= 26 {
				if upper == "--upper" {
					z01.PrintRune(rune(64 + customAtoi(numbers[i])))
				} else {
					z01.PrintRune(rune(96 + customAtoi(numbers[i])))
				}
			} else {
				z01.PrintRune(32)
			}
		}
	}
	z01.PrintRune('\n')
}
