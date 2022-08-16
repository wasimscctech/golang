package main

import "fmt"

func isleapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func main() {
	var year int
	fmt.Print("Enter the year: ")
	fmt.Scanln(&year)
	if isleapYear(year) {
		fmt.Print("The year is a Leap Year!")
	} else {
		fmt.Print("The year is not a Leap Year!")
	}
}
