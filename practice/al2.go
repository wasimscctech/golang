/*Write a GO program to get the absolute difference between n and 51. If n is greater than 51 return triple the absolute difference.
Sample Input:
53
30
51
Sample Output:
6
21
0*/

package main

import "fmt"

const Myval = 51

func difference(n int) int {
	if n > Myval {
		return (n - Myval) * 3
	} else {
		return (Myval - n)
	}
}

func main() {
	fmt.Println(difference(53))
	fmt.Println(difference(30))
	fmt.Println(difference(51))
}
