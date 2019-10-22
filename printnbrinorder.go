package piscine

import "github.com/01-edu/z01"

func size(n int) int {
	res := 0
	for ; n > 0; n /= 10 {
		res++
	}
	return res
}

func sortArr(arr []int) {
	change := false
	for i := 0; i < TabLen(table)-1; i++ {
		if table[i] > table[i+1] {
			temp := table[i]
			table[i] = table[i+1]
			table[i+1] = temp
			change = true
}

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0')
	}

	var array [19]int

	for i := 0; i <= 18; i++ {
		array[i] = 20
	}

	for i := 0; i < size(n); i++ {
		ten := 1
		for j := i; j > 0; j-- {
			ten = ten * 10
		}
		array[i] = (n / ten) % 10
	}
	sortArr(array[:])
	for i := 0; i < size(n); i++ {
		z01.PrintRune(rune(array[i] + 48))
	}
}
