package main

import "fmt"

func main() {
	fmt.Println("the adder result is: ", adder(2, 5))

	result1 := adder2(6, 5)
	fmt.Println("the result 2 is: ", result1)

	proresult := proadder(1, 2, 3, 4, 5)
	fmt.Println("the proresult is: ", proresult)

	proresult2, _ := proadder2(1, 2, 3, 4, 5, 6) //we are using underscore to avoid the input for string return type
	fmt.Println("the proresult2 is: ", proresult2)

	proresult3, myMessage := proadder2(1, 2, 3, 4, 5, 6, 10) //saving the returned string into a variable and printing it out
	fmt.Println("the proresult3 is: ", proresult3)
	fmt.Println("the myMessage is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	result := valOne + valTwo
	return result
}

func adder2(valone2, valTwo2 int) int {
	return valone2 + valTwo2
}

func proadder(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func proadder2(nums ...int) (int, string) {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum, "this is a string in return type"
}
