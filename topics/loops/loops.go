package main

import "fmt"

func main() {
	//golang has only one loop and that is a for loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("myRank: %v\n", i)
	} //for loop

	i := 5
	for i >= 1 {
		fmt.Printf("myRank: %v\n", i)
		i-- //while loop style and in reverse
	}
}
