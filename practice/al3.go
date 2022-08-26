/*
Write a GO program to check two given integers, and return true if one of them is 30 or if their sum is 30. Go to the editor
Sample Input:
30, 0
25, 5
20, 30
20, 25
Sample Output:
1
1
1
0*/
package main

import "fmt"

const n = 30

func is30(val1, val2 int) bool {
	if val1 == n || val2 == n || (val1+val2) == n {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(is30(30, 0))
	fmt.Println(is30(25, 5))
	fmt.Println(is30(20, 30))
	fmt.Println(is30(20, 25))
}
