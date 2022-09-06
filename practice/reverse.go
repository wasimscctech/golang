package main

import (
	"fmt"
)
func reverseValues(a, b string) (string, string) {
	return b, a //notice how multiple values are returned
}

func main() {
	val1, val2 := reverseValues("Wasim", "Shaikh") // notice how multiple values are assigned
	fmt.Println(val1, val2)
}

