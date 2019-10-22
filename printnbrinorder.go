package piscine

import "github.com/01-edu/z01"

func intToRune(digit int) rune {
	return rune(digit + 48)
}

func splitNumber(number int, arr []int) []int {
	if number == 0 {
		return arr
	}
	return splitNumber(number/10, append(arr, number%10))
}

func customLen(arr []int) int {
	count := 0
	for i := range arr {
		count += 1
		i++
	}
	return count
}

func customSort(arr []int) []int {
	for i := 0; i < customLen(arr); i++ {
		for j := i + 1; j < customLen(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
	} else {
		var arr []int
		arr = customSort(splitNumber(n, arr))
		for i := 0; i < customLen(arr); i++ {
			z01.PrintRune(intToRune(arr[i]))
		}
	}
}
