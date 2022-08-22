package main

import "fmt"

func main() {
	// f := func() {
	// 	fmt.Println("hello")
	// }
	// f()

	// f := func() string {
	// 	return "hello"
	// }
	// fmt.Println(f())

	func() {
		fmt.Println("hello")
	}()

}
