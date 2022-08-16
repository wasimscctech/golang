package main

import "fmt"

func main() {
	n1 := 0
	n2 := 1
	sum := 0
	for n3 := 0; n3 < 10; n3 = (n1 + n2) {
		n1 = n2
		n2 = n3
		fmt.Print(" ", n3)
		sum += n3
	}
	fmt.Printf("\nthe sum is %v", sum)
}
