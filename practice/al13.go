/*
Write a go program to check if one given temperatures is less than 0 and the other is greater than 100. Go to the editor
Sample Input:
120, -1
-1, 120
2, 120
Sample Output:
1
1
0*/

package main

import "fmt"

func isTempratureTrue(temp1, temp2 float64) bool {
	if temp1 > 100 && temp2 < 0 || temp1 < 0 && temp2 > 100 {
		return true
	}
	return false
}
func main() {
	fmt.Println(isTempratureTrue(120, -1))
	fmt.Println(isTempratureTrue(-1, 120))
	fmt.Println(isTempratureTrue(2, 120))
}
