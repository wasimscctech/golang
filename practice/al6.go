/*Write a Go program to remove the character in a given position of a given string. The given position will be in the range 0..string length -1 inclusive.
Sample Input:
"Python", 1
"Python", o
"Python", 4
Sample Output:
Pthon
ython
Pythn*/
package main

import "fmt"

func removemycharacter(myString string, index int) {
	fmt.Println((myString[:index] + myString[index+1:]))
}

func main() {
	removemycharacter("python", 1)
	removemycharacter("python", 0)
	removemycharacter("python", 4)

}
