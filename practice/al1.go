/* Write a Go program to compute the sum of the two given integer values. If the two values are the same, then return triple their sum. Go to the editor
Sample Input
1, 2
3, 2
2, 2
Sample Output:
3
5
12 */

package main

import "fmt"

func sum(val1, val2 int) int {
	return val1 + val2
}

func main() {

	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 2))
	fmt.Println(sum(2, 2))
}
