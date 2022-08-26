/*Write a GO program to create a new string where 'if' is added to the front of a given string. If the string already begins with 'if', return the string unchanged.
Sample Input:
"if else"
"else"
Sample Output:
if else
if else*/

package main

import "fmt"

func ifstring(myString string) {
	if myString[0:2] == "if" {
		fmt.Println(myString)
	} else {
		fmt.Println("if " + myString)
	}
}

func main() {
	ifstring("if else")
	ifstring("else")
	ifstring("Wasim")
}
