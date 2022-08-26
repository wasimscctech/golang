/*Write a go program to create a new string taking the first 3 characters of a given string and return the string with the 3 characters added at both the front and back. If the given string length is less than 3, use whatever characters are there. Go to the editor
Sample Input:
"Python"
"JS"
"Code"
Sample Output:
PytPythonPyt
JSJSJS
CodCodeCod*/

package main

import "fmt"

func last3Characters(myString string) string {
	if len(myString) >= 3 {
		subStr := myString[0:3]
		return subStr + myString + subStr
	}
	return myString + myString + myString
}

func main() {
	fmt.Println(last3Characters("Wasim"))
	fmt.Println(last3Characters("go"))
	fmt.Println(last3Characters("code"))
}
