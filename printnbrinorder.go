package piscine

import (
	"sort"

	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var nbrs []int

	for f := n; f != 0; f = (f / 10) {
		if f != 0 {
			k := (f % 10) + '0'
			nbrs = append(nbrs, k)
		} else {
			z01.PrintRune('0')
		}
	}

	sort.Sort(sort.IntSlice(nbrs))
	for _, v := range nbrs {
		z01.PrintRune(rune(v))
	}
}
