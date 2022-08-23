package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our services: ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("type of rating %T", input)
}
