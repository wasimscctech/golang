package main

import "fmt"

func main() {
	var result string
	score := 10
	if score > 5 && score < 10 {
		result = "good work!"
	} else if score == 10 {
		result = "exellent! Perfect score!"
	} else {
		result = "better luck next time!"
	}
	fmt.Println(result)

	if num := 4; num > 5 { //in go we can initialize the variable before after if and before conditional statement
		fmt.Println("the number is greater than 5")
	} else {
		fmt.Println("the is not greater than 5")
	}
}
