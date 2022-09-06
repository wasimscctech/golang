package main

import "fmt"

func linearSearch(myList []int, key int) bool {
	for _, item := range myList {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{13, 24, 54, 655, 875, 65, 23, 65, 87}
	// fmt.Println("the type is %T", items)
	fmt.Println(linearSearch(items, 6))
}
