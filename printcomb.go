package piscine

import "github.com/01-edu/z01"

func PrintComb() {
    for i:= '0'; i <= '9'; i++ {
           for j := i + 1; j <= '9'; j++ {
                   for k := j + 1; k <= '9'; k++ {
                           z01.PrintRune(i)
                           z01.PrintRune(j)
                           z01.PrintRune(k)
                           if i != 55 {
                                   z01.PrintRune(44)
                                   z01.PrintRune(32)
                           }
                   }
           }
    }
    z01.PrintRune(10)
}
