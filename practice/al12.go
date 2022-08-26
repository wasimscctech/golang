/*Write a go program to check if a given string starts with 'C#' or not. Go to the editor
Sample Input:
"C++ Sharp"
"C#"
"C++"
Sample Output:
1
1
0*/

package main

import "fmt"

func isStartswithGo(myString string) bool {
	if myString[:2] == "Go" {
		return true
	}
	return false
}
func main() {
	fmt.Println(isStartswithGo("Golang"))
	fmt.Println(isStartswithGo("Java"))
	fmt.Println(isStartswithGo("Go language"))
}
