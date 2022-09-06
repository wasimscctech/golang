package main

import "fmt"

func linearSearch1(myList []int, key int) bool {
	for _, item := range myList {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{12, 43, 54, 65, 76, 879, 98}
	fmt.Println(linearSearch1(items, 65))
}
