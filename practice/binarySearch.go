package main

import "fmt"

func binarySearch(items []int, key int) bool {
	low := 0
	high := len(items) - 1
	for low <= high {
		median := (low + high) / 2
		if items[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(items) || items[low] != key {
		return false
	}
	return true

}

func main() {
	items := []int{12, 23, 43, 54, 65, 76, 87, 98}
	fmt.Println(binarySearch(items, 76))
}
