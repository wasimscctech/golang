/*Write a go program to exchange the first and last characters in a given string and return the new string.
Sample Input:
"abcd"
"a"
"xy"
Sample output:
dbca
a
yx*/

package main

import "fmt"

func changeEndCharaters(myString string) {
	index := len(myString)
	NewString := myString[index-1:] + myString[1:index-1] + myString[0:1]
	fmt.Println(NewString)
}

func main() {
	changeEndCharaters("abcd")
	changeEndCharaters("xy")
}
