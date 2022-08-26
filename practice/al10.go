/*Write a C++ program to check if a given positive number is a multiple of 3 or a multiple of 7. Go to the editor
Sample Input
3
14
12
37
Sample Output:
1
1
1
0*/

package main

import "fmt"

func isMultiple(myVal int) bool {
	if myVal%3 == 0 || myVal%7 == 0 {
		return true
	}
	return false

}

func main() {
	fmt.Println(isMultiple(3))
	fmt.Println(isMultiple(14))
	fmt.Println(isMultiple(12))
	fmt.Println(isMultiple(37))
}
