package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	output := os.Args[0]

	for _, char := range output {
		z01.PrintRune(char)

	}

	z01.PrintRune('\n')
}
