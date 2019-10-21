
package piscine

func StrLen(str string) int {
	count := 0
	for _, char := range str {
		if string(char) != "" {
			count += 1
		}
	}
	return count
}
