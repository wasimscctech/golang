package main

import "fmt"

func main() {
	// defer keywork will help the code to get executed at the end. incase of multiple defer keywords: the first defer will get executed at the end

	defer fmt.Println("a code after the db connection closing")
	fmt.Println("Data base connection opening")
	defer fmt.Println("Data base connection closing")
	defer fmt.Println("a code that needs to be there before closing db connection closing")
	fmt.Println("Data base manipulation")
}
