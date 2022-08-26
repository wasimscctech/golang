/*Write a C++ program to compute the sum of the two given integers. If one of the given integer value is in the range 10..20 inclusive return 18. Go to the editor
Sample Input:
3, 7
10, 11
10, 20
21, 220
Sample Output:
10
18
18
241*/

package main

import "fmt"

func sum(num1, num2 int) int {
	if num1 >= 10 && num1 <= 20 || num2 >= 10 && num2 <= 20 {
		return 18
	}
	return num1 + num2
}

func main() {
	fmt.Println(sum(3, 7))
	fmt.Println(sum(10, 11))
	fmt.Println(sum(21, 220))
	fmt.Println(sum(3, 17))

}
