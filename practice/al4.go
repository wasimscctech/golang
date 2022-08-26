/*Write a GO program to check a given integer and return true if it is within 10 of 100 or 200. Go to the editor
Sample Input:
103
90
89
Sample Output:
1
1
0*/

package main

import (
	"fmt"
	"math"
)

func isWithin10(val float64) bool {
	if math.Abs((val-100)) <= 10 || math.Abs((val-200)) <= 10 {
		return true
	}
	return false
}
func main() {
	fmt.Println(isWithin10(103))
	fmt.Println(isWithin10(90))
	fmt.Println(isWithin10(89))
	fmt.Println(isWithin10(209))
	fmt.Println(isWithin10(20))
	fmt.Println(math.Abs((209 - 100)))
	fmt.Println(math.Abs((9 - 100)))

}
