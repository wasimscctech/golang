package main

import (
	"fmt"
)

func main() {
	// //separate declaration and initialization
	// empSal := make(map[string]int) //declaration
	// empSal = map[string]int{
	// 	"Wasim": 50000,
	// 	"Almas": 40000,
	// 	"aftab": 30000,
	// } //initialization
	// fmt.Println(empSal)

	// declaration with initialization
	empSal := map[string]int{
		"Wasim": 1000000,
		"almas": 900000,
		"aftab": 800000,
	}
	// fmt.Println(empSal)
	// fmt.Println(len(empSal))
	// fmt.Println(empSal["Wasim"]) //getting value of wasim
	// empSal["Wasim"] = 1500000    //reassigning the value of wasim
	// fmt.Println(empSal["Wasim"])
	// delete(empSal, "aftab") //deleting a key value pair from map
	// fmt.Println(empSal)

	fmt.Println(empSal["azim"]) //this will print zero as azim key is not present in the map. and this could be a trap so to avoid this we will check the availablity

	azim, ok := empSal["azim"]  //this a way to validate if the key is present in the map. ok can be anything.
	fmt.Println(azim, ok)       //we are getting the integer value of azim and "ok" will print the boolean value. we need to avoid the integer value as that is not neccessory and that will give us zero"0" for the key that is not present in the map. (o/p 0, false)
	_, flag := empSal["irshad"] //underscore will help us to skip the variable which we dont need
	fmt.Println(flag)           //it will print false as irshad is not present

	_, flag2 := empSal["Wasim"]
	fmt.Println(flag2) //o/p true
}
