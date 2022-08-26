/*Write a go program to check two given integers whether either of them is in the range 100..200 inclusive. Go to the editor
Sample Input:
100, 199
250, 300
105, 190
Sample Output:
1
0
1*/

package main

import "fmt"

func isNumberin100s(num1, num2 int) bool {
	if num1%100 == 0 || num2%100 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isNumberin100s(100, 199))
	fmt.Println(isNumberin100s(10, 200))
	fmt.Println(isNumberin100s(105, 199))
}
