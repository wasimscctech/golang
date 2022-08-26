package main

import "fmt"

func main() {
	defer fmt.Println("this is the last line because of defer")
	defer fmt.Println("this is the 2nd last line because of defer")
	fmt.Println("this is the first line")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //the execution will be in reverse because of defer
	}
}
