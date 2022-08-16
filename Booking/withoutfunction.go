package main

import (
	"fmt"
	"strings"
)

func main() {
	// ticketName := "Umrah" //below line is same
	var ticketName = "Umrah" //if we dont use this variable below then this will be an error same with the imported package
	const UmrahTickets = 50
	var RemainingTickets uint = 50

	// fmt.Print("Hello, and welcome to wasim's ticket booking app") this will not add a new line but Println will add a new line after the given statement
	// fmt.Println("Hello, and welcome to", ticketName, "booking app") //Go will automatically add a space before and after the variable name when we save the program

	// fmt.Printf("UmrahTickets is %T, RemainingTickets is %T and ticketName is %T", UmrahTickets, RemainingTickets, ticketName) //%T is used to print a type of any data

	fmt.Printf("Hello, and welcome to %v booking app\n", ticketName) //this won't add a new line
	fmt.Printf("We have total of %v tickets and %v are still available.", UmrahTickets, RemainingTickets)
	fmt.Println("Get your tickets booked here")

	// var booking = [50]string{} //empty array below is the other way
	// var bookings [50]string //array

	var bookings []string //a slice
	// bookings := []string{} //a slice
	for RemainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // userName = "Wasim"

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2 //below line is same
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			RemainingTickets = RemainingTickets - userTickets //operations between different integer types will be an error

			// bookings[0] = firstName + " " + lastName //adding elements in an array
			bookings = append(bookings, firstName+" "+lastName) //adding elements in a slice

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", RemainingTickets, ticketName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0]) //this loop ends when iterated overall elements if the bookings list, then going to the next statement of our program.
			}
			fmt.Printf("these are the firstnames of all the bookings %v\n", firstNames)

			if RemainingTickets == 0 {
				fmt.Println("All the ticket is booked out. Please comeback next month.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firstname or last name is not valid.")
			}
			if !isValidEmail {
				fmt.Println("rmail address you entered doesn't have @")
			}
			if !isValidTicketNumber {
				fmt.Println("ticket number is not valid")
			}
		}
	}
	// city := "London"
	// switch city {
	// case "New york":
	// 	//execute code for new york
	// case "Singapore":
	// 	//execute code for Singapore
	// case "London":
	// 	//execute code for London
	// case "Berlin":
	// 	//execute code for Berlin
	// case "pune", "mumbai": //pune or mumbai
	// 	//execute code for pune or mumbai
	// case "hyderabad":
	// 	//execute code for hyderabad
	// default:
	// 	fmt.Print("No valid city selected")
	// }
}
