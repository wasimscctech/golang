package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// a method for struct User
func (u User) GetStatus() {
	fmt.Println("the status of user is: ", u.Status)
}

func (u User) GetEmail() {
	u.Email = "Wasim@google.com"
	fmt.Println("the email of user is: ", u.Email)
}

func main() {
	wasim := User{"Wasim", "wasim@google.com", true, 26}
	fmt.Println("prints only the values: ", wasim)                        //o/p : {Wasim wasim@google.com true 26}
	fmt.Printf("prints only the values:  %v\n", wasim)                    // o/p : {Wasim wasim@google.com true 26}
	fmt.Printf("The complete details are here: %+v\n", wasim)             //o/p : {Name:Wasim Email:wasim@google.com Status:true Age:26}
	fmt.Printf("for user %v, the email is %v\n", wasim.Name, wasim.Email) // o/p : for user Wasim, the email is wasim@google.co
	wasim.GetStatus()
	wasim.GetEmail() //this won't change the original email id of the user. as we are not sending the details to memory location.
}
