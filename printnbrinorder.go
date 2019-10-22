package piscine

import "github.com/01-edu/z01"

func size(n int) int {
	res := 0
	for ; n > 0; n /= 10 {
		res++
	}
	return res
}

func getLen(arr []int) int {
	length := 0
	for i := range arr {
		length++
		i++
	}
	return length
}

func sortArr(arr []int) {
	for i := 0; i < getLen(arr); i++ {
		for j := i + 1; j < getLen(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
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
