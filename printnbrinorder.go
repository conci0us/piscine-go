package piscine

import "github.com/01-edu/z01"

func size(n int) int {
	res := 0
	for ; n > 0; n /= 10 {
		res++
	}
	return res
}

func tabLen(table []int) int {
	length := 0
	for range table {
		length++
	}
	return length
}

func sortArr(arr []int) {
	for i := 0; i < tabLen(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
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
