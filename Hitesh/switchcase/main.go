package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //the rand should be from "math/rand" not from crypto
	diceNumber := rand.Intn(6) + 1
	fmt.Println("the value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the value is 1, now you can open or move 1 step")
	case 2:
		fmt.Println("the value is 2, you can move 2 step")
	case 3:
		fmt.Println("the value is 3, you can move 3 step")
	case 4:
		fmt.Println("the value is 4, you can move 4 step")
		fallthrough //this will print the next case if the current case is true
	case 5:
		fmt.Println("the value is 5, you can move 5 step")
	case 6:
		fmt.Println("the value is 6, you can move 6 step and throw another dice")
	default:
		fmt.Println("What was that!")
	}
}
