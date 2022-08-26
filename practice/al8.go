/*Write a go program to create a new string which is 4 copies of the 2 front characters of a given string. If the given string length is less than 2 return the original string. Go to the editor
Sample Input:
"C Sharp"
"JS"
"a"
Sample Output:
C C C C
JSJSJSJS
a*/
package main

import "fmt"

func return4Copies(myString string) {
	if len(myString) >= 2 {
		newString := myString[0:2]
		fmt.Println(newString + newString + newString + newString)
	} else {
		fmt.Println(myString)
	}
}

func main() {
	return4Copies("Go")
	return4Copies("C Sharp")
	return4Copies("W")

}
