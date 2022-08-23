package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	wasim := User{"Wasim", "wasim@google.com", true, 26}
	fmt.Println("prints only the values: ", wasim)
	fmt.Printf("prints only the values:  %v\n", wasim)
	fmt.Printf("The complete details are here: %+v\n", wasim)
	fmt.Printf("for user %v, the email is %v\n", wasim.Name, wasim.Email)

}
