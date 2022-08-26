/*Write a go program to create a new string with the last char added at the front and back of a given string of length 1 or more. Go to the editor
Sample Input:
"Red"
"Green"
"1"
Sample Output:
dRedd
nGreenn
111*/

package main

import "fmt"

func firstLastCharacterString(myString string) string {
	n := len(myString)
	endChar := myString[n-1:]
	return endChar + myString + endChar
}
func main() {
	fmt.Println(firstLastCharacterString("Red"))
	fmt.Println(firstLastCharacterString("Green"))
	fmt.Println(firstLastCharacterString("1"))
}
